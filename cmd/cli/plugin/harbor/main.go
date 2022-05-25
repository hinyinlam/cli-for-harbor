/*
 * CLI for Harbor
 * Copyright 2022 VMware, Inc.
 *
 * This product is licensed to you under the Apache 2.0 license (the "License").  You may not use this product except in compliance with the Apache 2.0 License.
 *
 * This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
 */

/*
 * Author: Hin Lam <hinl@vmware.com>
 */

package main

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"

	"github.com/aunum/log"
	"github.com/hinyinlam/cli-for-harbor/go/harbor/cmd"

	cliv1alpha1 "github.com/vmware-tanzu/tanzu-framework/apis/cli/v1alpha1"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/cli/command/plugin"
)

var descriptor = cliv1alpha1.PluginDescriptor{
	Name:        "harbor",
	Description: "*unofficial* harbor cli - alpha - Author: Hin Lam <hinl@vmware.com>",
	Version:     "v0.0.1",
	Group:       cliv1alpha1.ManageCmdGroup, // set group
}

func main() {
	cmd.AuthFile = filepath.Join(xdg.Home, ".config", "tanzu", "harbor", "auth.yaml")
	cmd.InitHarborApi()
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err)
	}
	p.AddCommands(
		cmd.RootCmd.Commands()...,
	)
	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}
