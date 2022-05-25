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
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/hinyinlam/cli-for-harbor/go/harbor/cmd"
)

func main() {
	cmd.AuthFile = filepath.Join(xdg.Home, ".config", "harbor", "auth.yaml")
	cmd.InitHarborApi()
	cmd.Execute()
}
