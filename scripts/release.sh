#!/usr/bin/env bash

# Exit on error. Append "|| true" if you expect an error.
set -o errexit
# Exit on error inside any functions or subshells.
set -o errtrace
# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset
# Catch the error in case mysqldump fails (but gzip succeeds) in `mysqldump | gzip`
set -o pipefail

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${__dir}/libs/logs.sh"
source "${__dir}/libs/git.sh"

# These environment variables can be used to alter the behavior of the release script.

DRY_RUN=${DRY_RUN:-"yes"} # Some kind of fail-safe to ensure that we're only pushing something when we mean it.

PROMQL_BUILDER_REPO=${PROMQL_BUILDER_REPO:-'git@github.com:grafana/promql-builder.git'}

CLEANUP_WORKSPACE=${CLEANUP_WORKSPACE:-"yes"} # Should the workspace be deleted after the script runs?
WORKSPACE_PATH=${WORKSPACE_PATH:-'./workspace'}

#################
### Usage ###
#################

# LOG_LEVEL=7 ./scripts/release.sh

#################
### Utilities ###
#################

function run_when_safe() {
  local command=${1}
  shift

  if [ "${DRY_RUN}" == "no" ]; then
    ${command} "$@"
  else
    warning "Dry run enabled: skipping execution of \"${command} $*\""
    info "Run this script with DRY_RUN=no to disable dry-run mode."
  fi
}

function clone_promql_builder() {
  local clone_into_dir="${1}"
  shift

  git clone "${PROMQL_BUILDER_REPO}" "${clone_into_dir}"
}

function should_abort_release() {
  local release_marker=${1}
  shift

  # No release file means no release was ever made: nothing to release.
  if [ ! -f "${release_marker}" ]; then
    return 0
  fi

  latest_tag=$(git_run "${promql_builder_path}" describe --tags --match 'v*.*.*' --abbrev=0 2>/dev/null || echo 'v0.0.0')
  current_marker=$(cat "${release_marker}")

  # The release marker and latest tags are equal: we just merged the
  # release PR and already performed the release. Let's prepare a new one.
  if [ "${latest_tag}" == "${current_marker}" ]; then
    return 0
  fi

  # The release marker and latest tags are different: we just merged the
  # release PR and should perform the release now.
  return 1
}

############
### Main ###
############

promql_builder_path="${WORKSPACE_PATH}/promql-builder"
release_file_marker=".release/tag"

if [ ! -f "${release_file_marker}" ]; then
  notice "Release marker not found, aborting."
    echo "is_release=no" >> "$GITHUB_OUTPUT"
  exit 0
fi

if [ "${DRY_RUN}" == "no" ]; then
  warning "Dry-run is OFF."
else
  notice "Dry-run is ON."
fi

debug "workspace path: ${WORKSPACE_PATH}"

# Just in case there are leftovers from a previous run.
rm -rf "${WORKSPACE_PATH}"

info "Cloning promql-builder into ${promql_builder_path}"
clone_promql_builder "${promql_builder_path}"

if should_abort_release "${release_file_marker}"; then
    notice "No release is pending, aborting."
    echo "is_release=no" >> "$GITHUB_OUTPUT"
    exit 0
fi

next_tag=$(cat "${promql_builder_path}/${release_file_marker}")
run_when_safe git_run "${promql_builder_path}" tag "${next_tag}"
# Go needs a specific tag since we declared the module in a subdirectory
# See https://go.dev/ref/mod#vcs-version
run_when_safe git_run "${promql_builder_path}" tag "go/${next_tag}" 
run_when_safe git_run "${promql_builder_path}" push origin --tags

echo "is_release=yes" >> "$GITHUB_OUTPUT"