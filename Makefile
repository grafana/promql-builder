COG_VERSION = v0.1.13
COG_DIR     = $(shell go env GOPATH)/bin/cog-$(COG_VERSION)
COG_BIN     = $(COG_DIR)/cli

# Within devbox
ifneq "$(DEVBOX_CONFIG_DIR)" ""
    RUN_DEVBOX:=
else # Normal shell
    RUN_DEVBOX:=devbox run
endif

.PHONY: install-cog
install-cog: $(COG_BIN)

$(COG_BIN):
	@echo "Installing Cog version $(COG_VERSION)"
	@mkdir -p $(COG_DIR)
	GOBIN=$(COG_DIR) go install github.com/grafana/cog/cmd/cli@$(COG_VERSION)
	@touch $@

.PHONY: validate-config
validate-config: install-cog
	$(COG_BIN) inspect --config .cog/config.yaml --ir builders

.PHONY: release
release: install-cog
	COG_CMD=$(COG_BIN) $(RUN_DEVBOX) ./scripts/release.sh

.PHONY: docs
docs:
	$(RUN_DEVBOX) ./scripts/build-docs.sh

.PHONY: serve-docs
serve-docs:
	SERVE=please $(RUN_DEVBOX) ./scripts/build-docs.sh
