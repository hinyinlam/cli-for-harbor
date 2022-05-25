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
BUILD_VERSION=v0.0.1
LD_FLAGS = -X 'github.com/vmware-tanzu/tanzu-framework/pkg/v1/cli/common.IsContextAwareDiscoveryEnabled=false'
ENVS ?= linux-amd64 windows-amd64 darwin-amd64

.PHONY: all
all: update_go_dependencies package_standalone package_tanzu_plugin

.PHONY: update_go_dependencies
update_go_dependencies:
	go get -u github.com/goharbor/go-client

.PHONY: clean
clean: tanzu_plugin_delete tanzu_plugin_delete_discovery
	#- cd cmd && rm gen_*.go
	- rm -rf standalone_binary/* artifacts/*
	- rm ~/Library/Preferences/harbor/auth.yaml
	- rm -rf ~/.config/tanzu-plugins/*/*/*/harbor
	- rm tanzu-framework-plugins-harbor-allarch.tar.gz  harbor-standalone.tar.gz
	#- rm -rf /tmp/cobra-codegen
	- rm -rf ./dist/
	- rm -rf ~/.config/harbor/ ~/.config/tanzu/harbor/


.PHONY: standalone_binary
standalone_binary:
	for os in darwin linux windows; do \
		for arch in amd64 arm64 ; do \
			echo env GOOS=$$os GOARCH=$$arch go build -o standalone_binary/harbor-$$os\_$$arch cmd/cli/standalone.go; \
			env GOOS=$$os GOARCH=$$arch go build -o standalone_binary/harbor-$$os\_$$arch cmd/cli/standalone.go; \
		done \
	done
		

.PHONY: tanzu_plugin
tanzu_plugin: 
	for OS in darwin linux windows ; do \
		for ARCH in amd64 ; do \
			tanzu builder cli compile --version v0.0.1 --ldflags "$(LD_FLAGS) -X 'github.com/vmware-tanzu/tanzu-framework/pkg/v1/config.DefaultStandaloneDiscoveryType=oci'" --tags "HinLamMakefileBuild" --path ./cmd/cli/plugin/ --artifacts artifacts/$$OS/$$ARCH/cli --target $$OS\_$$ARCH ; \
		done \
	done

.PHONY: tanzu_plugin_delete
tanzu_plugin_delete:
	- tanzu plugin delete harbor
	- rm -r ~/.config/tanzu/harbor

.PHONY: tanzu_plugin_install_discovery
tanzu_plugin_install_discovery:
	tanzu plugin source add --name harbor-local-source --type local --uri ${HOME}/.config/tanzu-plugins/discovery/harbor
	tanzu plugin install harbor

.PHONY: tanzu_plugin_delete_discovery
tanzu_plugin_delete_discovery:
	- tanzu plugin source delete harbor-local-source

.PHONY: package_tanzu_plugin
package_tanzu_plugin: tanzu_plugin
	tar jcvf tanzu-framework-plugins-harbor-allarch.tar.gz artifacts Makefile

.PHONY: package_standalone
package_standalone: standalone_binary
	tar jcvf harbor-standalone.tar.gz standalone_binary Makefile

.PHONY: uninstall_tanzu_cli
uninstall_tanzu_cli:
	- sudo rm -rf /usr/local/bin/tanzu ~/tanzu/cli
	- rm -rf ~/Library/Application\ Support/tanzu
	- rm -rf ~/Library/Application\ Support/tanzu-cli
	- rm -rf ~/.config/tanzu ~/.config/tanzu-plugin
	- rm -rf ~/.cache/tanzu
	- rm -rf ~/.tanzu

OCI_REGSITRY=demo.goharbor.io
.PHONY: publish_to_oci
publish_to_oci:
	tanzu builder publish --type oci --plugins harbor --version v0.0.1 --os-arch "$(ENVS)" --oci-discovery-image $(OCI_REGSITRY)/tanzu-plugins/discovery/harbor:v0.0.1 --oci-distribution-image-repository $(OCI_REGSITRY)/tanzu-plugins/distribution/ --input-artifact-dir ./artifacts

.PHONY: add_oci_harbor_plugin_source
add_oci_harbor_plugin_source:
	echo This command does not work
	tanzu plugin source delete hintest && tanzu plugin source add  --name hintest --type oci --uri demo.goharbor.io/tanzu-plugins/discovery/harbor:v0.0.1


XDG_CONFIG_HOME := ${HOME}/.config
TANZU_PLUGIN_PUBLISH_PATH ?= $(XDG_CONFIG_HOME)/tanzu-plugins
.PHONY: publish_to_local
publish_to_local:
	tanzu builder publish --type local --plugins "harbor" --version v0.0.1 --os-arch "$(ENVS)" --local-output-discovery-dir "$(TANZU_PLUGIN_PUBLISH_PATH)/discovery/harbor" --local-output-distribution-dir "$(TANZU_PLUGIN_PUBLISH_PATH)/distribution" --input-artifact-dir ./artifacts

