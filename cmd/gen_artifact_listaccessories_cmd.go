/*
 * CLI for Harbor
 * Copyright 2022 VMware, Inc.
 *
 * This product is licensed to you under the Apache 2.0 license (the "License").  You may not use this product except in compliance with the Apache 2.0 License.
 *
 * This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

/*
 * This is a generated file by Cobra-codegen (https://github.com/hinyinlam/cli-for-harbor/cobra-codegen) Do NOT edit manually
 * Author: Hin Lam <hinl@vmware.com>
 */

package cmd

import (
	"fmt"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/artifact"
	"github.com/spf13/cobra"
	"os"
	"gopkg.in/yaml.v2"
)

var artifactListAccessoriesCmd= &cobra.Command{
	Use:   "listaccessories",
	Short: "listaccessories artifact",
	Long:  `Sub-command for ListAccessories`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&ListAccessoriesParamsWrapperDefaultValue)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		
			fmt.Println("show-rquest-yaml done")
			os.Exit(0)
		}
		
		listOfCmdParams := map[string]string{
		
			
		
			
			"Page": "*int64",
			
			"PageSize": "*int64",
			
			"ProjectName": "string",
			
			"Q": "*string",
			
			"Reference": "string",
			
			"RepositoryName": "string",
			
			"Sort": "*string",
			
			"timeout": "time.Duration",
			
		
		}

		
		localListAccessoriesParams := &ListAccessoriesParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localListAccessoriesParams, err = UnmarshalFileFlag(cmd, localListAccessoriesParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localListAccessoriesParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("artifact") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*artifact.Client).ListAccessories(cliContext, &localListAccessoriesParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	artifactListAccessoriesCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	artifactListAccessoriesCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	artifactListAccessoriesCmd.Flags().StringP("Page", "", "", "Type: *int64")
		
	artifactListAccessoriesCmd.Flags().StringP("PageSize", "", "", "Type: *int64")
		
	artifactListAccessoriesCmd.Flags().StringP("ProjectName", "", "", "Type: string")
		
	artifactListAccessoriesCmd.Flags().StringP("Q", "", "", "Type: *string")
		
	artifactListAccessoriesCmd.Flags().StringP("Reference", "", "", "Type: string")
		
	artifactListAccessoriesCmd.Flags().StringP("RepositoryName", "", "", "Type: string")
		
	artifactListAccessoriesCmd.Flags().StringP("Sort", "", "", "Type: *string")
		
	artifactListAccessoriesCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	artifactCmd.AddCommand(artifactListAccessoriesCmd)

}


type ListAccessoriesParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       artifact.ListAccessoriesParams	`yaml:"spec" json:"spec"`
}
var ListAccessoriesParamsWrapperDefaultValue = ListAccessoriesParamsWrapper{
	Kind: "artifact/ListAccessoriesParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: artifact.ListAccessoriesParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

