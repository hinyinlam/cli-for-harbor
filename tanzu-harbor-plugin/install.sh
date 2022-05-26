#!/bin/bash
SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]:-$0}"; )" &> /dev/null && pwd 2> /dev/null; )";
TANZU_PLUGIN_PUBLISH_PATH=$HOME/.config/tanzu-plugins

cp -rvf $SCRIPT_DIR/distribution/* $TANZU_PLUGIN_PUBLISH_PATH/distribution/
cp -rvf $SCRIPT_DIR/discovery/* $TANZU_PLUGIN_PUBLISH_PATH/discovery/
tanzu plugin source add --name harbor-local-source --type local --uri $TANZU_PLUGIN_PUBLISH_PATH/discovery/harbor
tanzu plugin install harbor