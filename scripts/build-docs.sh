#!/usr/bin/env bash

# Exit on error. Append "|| true" if you expect an error.
set -o errexit
# Exit on error inside any functions or subshells.
set -o errtrace
# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset
# Catch the error in case mysqldump fails (but gzip succeeds) in `mysqldump |gzip`
set -o pipefail

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "${__dir}/libs/logs.sh"

SERVE=${SERVE:-"no"}

LANGUAGES="go;python;typescript"

build_dir="${__dir}/../build"
docs_build_dir="${build_dir}/docs"
mkdocs_dir="${__dir}/../.mkdocs"

rm -rf "${build_dir}"
mkdir -p "${build_dir}"

cp -R "${mkdocs_dir}" "${docs_build_dir}"

for language in ${LANGUAGES//;/ } ; do
    echo "🪧 Preparing documentation for language ${language}"
    cp -R ${language}/docs/* "${docs_build_dir}/docs/${language}"
done

echo "🪧 Building documentation website"

if [ "${SERVE}" == "no" ]; then
    zensical build -f ${docs_build_dir}/mkdocs.yml 
else
    zensical serve -f ${docs_build_dir}/mkdocs.yml 
fi
