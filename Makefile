# /*
# * CLI for Harbor
# * Copyright 2022 VMware, Inc.
# *
# * This product is licensed to you under the Apache 2.0 license (the "License").  You may not use this product except in compliance with the Apache 2.0 License.
# *
# * This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
# */
# /*
# * Author: Hin Lam <hinl@vmware.com>
# *

# This file is for both building and installing Harbor command in both standalone and `tanzu plugin` form factor
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BUILD_VERSION=v0.0.4
LD_FLAGS = -X 'github.com/vmware-tanzu/tanzu-framework/pkg/v1/cli/common.IsContextAwareDiscoveryEnabled=false'


### Standalone Harbor CLI Section ###
.PHONY: standalone_binary
standalone_binary:
	for os in darwin linux windows; do \
		for arch in amd64 arm64 ; do \
			echo env GOOS=$$os GOARCH=$$arch go build -o standalone_binary/harbor-$$os\_$$arch cmd/cli/standalone.go; \
			env GOOS=$$os GOARCH=$$arch go build -o standalone_binary/harbor-$$os\_$$arch cmd/cli/standalone.go; \
		done \
	done

.PHONY: package_standalone
package_standalone: standalone_binary
	tar jcvf harbor-standalone.tar.gz standalone_binary Makefile

.PHONY: standalone_binary_clean
standalone_binary_clean:
	- rm -rf standalone_binary/* artifacts/*
	- rm ~/Library/Preferences/harbor/auth.yaml
	- rm -rf ~/.config/tanzu-plugins/*/*/*/harbor
	- rm tanzu-framework-plugins-harbor-allarch.tar.gz  harbor-standalone.tar.gz
		
### Tanzu Plugin Section ###
.PHONY: tanzu_plugin_clean
tanzu_plugin_clean: tanzu_plugin_delete tanzu_plugin_delete_discovery tanzu_plugin_non_builder_clean tanzu_plugin_using_builder_clean
	- rm -f tanzu-harbor-plugin/release-version.txt

.PHONY: package_tanzu_plugin
#package_tanzu_plugin: tanzu_plugin_using_builder tanzu_plugin_non_builder
package_tanzu_plugin: tanzu_plugin_non_builder
	tar jcvf tanzu-harbor-plugin-allarch.tar.gz tanzu-harbor-plugin

RED='\033[0;31m'
NC='\033[0m'
.PHONY: tanzu_plugin_using_builder
tanzu_plugin_using_builder: 
	echo -e "$(RED) ### This 'tanzu builder' plugin currently generate incorrect layout, for demo purpose ### $(NC)"; \
	for OS in darwin linux windows ; do \
		for ARCH in amd64 ; do \
			tanzu builder cli compile --version $(BUILD_VERSION) --ldflags "$(LD_FLAGS) -X 'github.com/vmware-tanzu/tanzu-framework/pkg/v1/config.DefaultStandaloneDiscoveryType=oci'" --tags "HinLamMakefileBuild" --path ./cmd/cli/plugin/ --artifacts artifacts/$$OS/$$ARCH/cli --target $$OS\_$$ARCH ; \
		done \
	done

.PHONY: tanzu_plugin_using_builder_clean
tanzu_plugin_using_builder_clean:
	- rm -rf ./artifacts/*

TANZU_HARBOR_PLUGIN_DISTRIBUTION_DIR=tanzu-harbor-plugin/distribution
.PHONY: tanzu_plugin_non_builder
tanzu_plugin_non_builder:
	for os in darwin linux windows; do \
		for arch in amd64 arm64 ; do \
			export output_folder=$(TANZU_HARBOR_PLUGIN_DISTRIBUTION_DIR)/$$os/$$arch/harbor/$(BUILD_VERSION); \
			mkdir -p $$output_folder; \
			echo env GOOS=$$os GOARCH=$$arch go build -o $$output_folder/tanzu-harbor-$$os\_$$arch cmd/cli/plugin/harbor/main.go; \
			env GOOS=$$os GOARCH=$$arch go build -o $$output_folder/tanzu-harbor-$$os\_$$arch cmd/cli/plugin/harbor/main.go; \
		done \
	done

.PHONY: tanzu_plugin_non_builder_clean
tanzu_plugin_non_builder_clean:
	for os in darwin linux windows; do \
		for arch in amd64 arm64 ; do \
			export output_folder=$(TANZU_HARBOR_PLUGIN_DISTRIBUTION_DIR)/$$os/$$arch/harbor/$(BUILD_VERSION); \
			rm -rf $$output_folder/tanzu-* ; \
		done \
	done


.PHONY: tanzu_plugin_delete
tanzu_plugin_delete:
	- tanzu plugin delete harbor
	- rm -r ~/.config/tanzu/harbor

XDG_CONFIG_HOME := ${HOME}/.config
TANZU_PLUGIN_PUBLISH_PATH ?= $(XDG_CONFIG_HOME)/tanzu-plugins
.PHONY: tanzu_plugin_install_discovery
tanzu_plugin_install_discovery:
	cp -rvf $(ROOT_DIR)/tanzu-harbor-plugin/distribution/* $(TANZU_PLUGIN_PUBLISH_PATH)/distribution/
	cp -rvf $(ROOT_DIR)/tanzu-harbor-plugin/discovery/* $(TANZU_PLUGIN_PUBLISH_PATH)/discovery/
	tanzu plugin source add --name harbor-local-source --type local --uri $(TANZU_PLUGIN_PUBLISH_PATH)/discovery/harbor
	tanzu plugin install harbor

.PHONY: tanzu_plugin_delete_discovery
tanzu_plugin_delete_discovery:
	- tanzu plugin source delete harbor-local-source

.PHONY: tanzu_plugin_release
tanzu_plugin_release: 
	git tag -l --sort=version:refname > tanzu-harbor-plugin/release-version.txt
	gh release upload `cat tanzu-harbor-plugin/release-version.txt` tanzu-harbor-plugin-allarch.tar.gz --clobber

### Tanzu CLI related ###
.PHONY: uninstall_tanzu_cli
uninstall_tanzu_cli:
	- sudo rm -rf /usr/local/bin/tanzu ~/tanzu/cli
	- rm -rf ~/Library/Application\ Support/tanzu
	- rm -rf ~/Library/Application\ Support/tanzu-cli
	- rm -rf ~/.config/tanzu ~/.config/tanzu-plugin
	- rm -rf ~/.cache/tanzu
	- rm -rf ~/.tanzu

### Common target ###
.PHONY: all
all: update_go_dependencies package_standalone package_tanzu_plugin

.PHONY: update_go_dependencies
update_go_dependencies:
	go get -u github.com/goharbor/go-client

.PHONY: clean
clean: standalone_binary_clean tanzu_plugin_clean gorelease_clean

### Releaser ###
.PHONY: release
release: gorelease_standalone tanzu_plugin_release

### GoReleaser ###
.PHONY: gorelease_standalone
gorelease_standalone:
	goreleaser --rm-dist 

.PHONY: gorelease_clean
gorelease_clean:
	- rm -rf ./dist/*

### GitHub release for tanzu-plugin ###

# Disable due to not working
# OCI_REGSITRY=demo.goharbor.io
# .PHONY: publish_to_oci
# publish_to_oci:
# 	tanzu builder publish --type oci --plugins harbor --version v0.0.1 --os-arch "$(ENVS)" --oci-discovery-image $(OCI_REGSITRY)/tanzu-plugins/discovery/harbor:v0.0.1 --oci-distribution-image-repository $(OCI_REGSITRY)/tanzu-plugins/distribution/ --input-artifact-dir ./artifacts

# Disable due to not working
# .PHONY: add_oci_harbor_plugin_source
# add_oci_harbor_plugin_source:
# 	echo This command does not work
# 	tanzu plugin source delete hintest && tanzu plugin source add  --name hintest --type oci --uri demo.goharbor.io/tanzu-plugins/discovery/harbor:v0.0.1

# Disable due to not working
# ENVS ?= linux-amd64 windows-amd64 darwin-amd64
# .PHONY: publish_to_local
# publish_to_local:
# 	tanzu builder publish --type local --plugins "harbor" --version v0.0.1 --os-arch "$(ENVS)" --local-output-discovery-dir "$(TANZU_PLUGIN_PUBLISH_PATH)/discovery/harbor" --local-output-distribution-dir "$(TANZU_PLUGIN_PUBLISH_PATH)/distribution" --input-artifact-dir ./artifacts

