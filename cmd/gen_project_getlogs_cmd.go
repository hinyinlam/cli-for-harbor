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
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/spf13/cobra"
	"os"
	"gopkg.in/yaml.v2"
)

var projectGetLogsCmd= &cobra.Command{
	Use:   "getlogs",
	Short: "getlogs project",
	Long:  `Sub-command for GetLogs`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&GetLogsParamsWrapperDefaultValue)
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
			
			"Sort": "*string",
			
			"timeout": "time.Duration",
			
		
		}

		
		localGetLogsParams := &GetLogsParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localGetLogsParams, err = UnmarshalFileFlag(cmd, localGetLogsParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localGetLogsParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("project") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*project.Client).GetLogs(cliContext, &localGetLogsParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	projectGetLogsCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	projectGetLogsCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	projectGetLogsCmd.Flags().StringP("Page", "", "", "Type: *int64")
		
	projectGetLogsCmd.Flags().StringP("PageSize", "", "", "Type: *int64")
		
	projectGetLogsCmd.Flags().StringP("ProjectName", "", "", "Type: string")
		
	projectGetLogsCmd.Flags().StringP("Q", "", "", "Type: *string")
		
	projectGetLogsCmd.Flags().StringP("Sort", "", "", "Type: *string")
		
	projectGetLogsCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	projectCmd.AddCommand(projectGetLogsCmd)

}


type GetLogsParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       project.GetLogsParams	`yaml:"spec" json:"spec"`
}
var GetLogsParamsWrapperDefaultValue = GetLogsParamsWrapper{
	Kind: "project/GetLogsParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: project.GetLogsParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

