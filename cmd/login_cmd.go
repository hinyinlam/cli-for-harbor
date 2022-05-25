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

package cmd

import (
	"fmt"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/health"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Harbor registry",
	Long:  `Authenticate with Harbor Registry`,
	Run: func(cmd *cobra.Command, args []string) {
		localClientSetConfigParam := &ClientSetConfigWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename != "" {
			localClientSetConfigParam, _ = UnmarshalFileFlag(cmd, localClientSetConfigParam)
		} else {
			fitCmdParamIntoDataStruct(cmd, map[string]string{"url": "string", "username": "string", "password": "string"}, &localClientSetConfigParam.Spec)
			if localClientSetConfigParam.Spec.Username == "" {
				fmt.Println("please specific --username")
				os.Exit(-1)
			}
			if localClientSetConfigParam.Spec.Password == "" {
				fmt.Println("please specific --password")
				os.Exit(-1)
			}

			if localClientSetConfigParam.Spec.URL == "" {
				fmt.Println("please specific --url")
				os.Exit(-1)
			}
			if !(strings.HasPrefix(localClientSetConfigParam.Spec.URL, "http://") ||
				strings.HasPrefix(localClientSetConfigParam.Spec.URL, "https://")) {
				fmt.Printf("--url is incorrect - URL should begin with https:// or http://\n")
				os.Exit(-1)
			}
			_, err = url.ParseRequestURI(localClientSetConfigParam.Spec.URL)
			if err != nil {
				fmt.Printf("--url is incorrect %+v\n", err)
				os.Exit(-1)
			}
		}

		clientSetConfig = localClientSetConfigParam.Spec
		fmt.Println("Login using:")
		fmt.Println("API URL: ", clientSetConfig.URL)
		fmt.Println("Username: ", clientSetConfig.Username)
		fmt.Println("SSL: ", strconv.FormatBool(!clientSetConfig.Insecure))
		InitHarborApi()
		_, err = harborApi.Health.GetHealth(cliContext, &health.GetHealthParams{})
		if err != nil {
			panic(err)
		}
		SaveAuthInfo(localClientSetConfigParam.Spec)
		fmt.Println("Login success")
	},
}

func init() {
	loginCmd.Flags().StringP("filename", "f", "", "filename to request parameter")
	loginCmd.Flags().StringP("url", "", "", "API URL, eg: https://myharbor.somedomain.local")
	loginCmd.Flags().StringP("username", "", "", "username")
	loginCmd.Flags().StringP("password", "", "", "password")
	RootCmd.AddCommand(loginCmd)
}
