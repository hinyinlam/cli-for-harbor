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
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/label"
	"github.com/spf13/cobra"
	"os"
	"gopkg.in/yaml.v2"
)

var labelListCmd= &cobra.Command{
	Use:   "list",
	Short: "list label",
	Long:  `Sub-command for List`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&ListLabelsParamsWrapperDefaultValue)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		
			fmt.Println("show-rquest-yaml done")
			os.Exit(0)
		}
		
		listOfCmdParams := map[string]string{
		
			
		
			
			"Name": "*string",
			
			"Page": "*int64",
			
			"PageSize": "*int64",
			
			"ProjectID": "*int64",
			
			"Q": "*string",
			
			"Scope": "*string",
			
			"Sort": "*string",
			
			"timeout": "time.Duration",
			
		
		}

		
		localListLabelsParams := &ListLabelsParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localListLabelsParams, err = UnmarshalFileFlag(cmd, localListLabelsParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localListLabelsParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("label") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*label.Client).ListLabels(cliContext, &localListLabelsParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	labelListCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	labelListCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	labelListCmd.Flags().StringP("Name", "", "", "Type: *string")
		
	labelListCmd.Flags().StringP("Page", "", "", "Type: *int64")
		
	labelListCmd.Flags().StringP("PageSize", "", "", "Type: *int64")
		
	labelListCmd.Flags().StringP("ProjectID", "", "", "Type: *int64")
		
	labelListCmd.Flags().StringP("Q", "", "", "Type: *string")
		
	labelListCmd.Flags().StringP("Scope", "", "", "Type: *string")
		
	labelListCmd.Flags().StringP("Sort", "", "", "Type: *string")
		
	labelListCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	labelCmd.AddCommand(labelListCmd)

}


type ListLabelsParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       label.ListLabelsParams	`yaml:"spec" json:"spec"`
}
var ListLabelsParamsWrapperDefaultValue = ListLabelsParamsWrapper{
	Kind: "label/ListLabelsParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: label.ListLabelsParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

