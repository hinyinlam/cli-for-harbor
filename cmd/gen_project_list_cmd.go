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

var projectListCmd= &cobra.Command{
	Use:   "list",
	Short: "list project",
	Long:  `Sub-command for List`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&ListProjectsParamsWrapperDefaultValue)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		
			fmt.Println("show-rquest-yaml done")
			os.Exit(0)
		}
		
		listOfCmdParams := map[string]string{
		
			
		
			
			"Name": "*string",
			
			"Owner": "*string",
			
			"Page": "*int64",
			
			"PageSize": "*int64",
			
			"Public": "*bool",
			
			"Q": "*string",
			
			"Sort": "*string",
			
			"WithDetail": "*bool",
			
			"timeout": "time.Duration",
			
		
		}

		
		localListProjectsParams := &ListProjectsParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localListProjectsParams, err = UnmarshalFileFlag(cmd, localListProjectsParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localListProjectsParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("project") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*project.Client).ListProjects(cliContext, &localListProjectsParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	projectListCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	projectListCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	projectListCmd.Flags().StringP("Name", "", "", "Type: *string")
		
	projectListCmd.Flags().StringP("Owner", "", "", "Type: *string")
		
	projectListCmd.Flags().StringP("Page", "", "", "Type: *int64")
		
	projectListCmd.Flags().StringP("PageSize", "", "", "Type: *int64")
		
	projectListCmd.Flags().StringP("Public", "", "", "Type: *bool")
		
	projectListCmd.Flags().StringP("Q", "", "", "Type: *string")
		
	projectListCmd.Flags().StringP("Sort", "", "", "Type: *string")
		
	projectListCmd.Flags().StringP("WithDetail", "", "", "Type: *bool")
		
	projectListCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	projectCmd.AddCommand(projectListCmd)

}


type ListProjectsParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       project.ListProjectsParams	`yaml:"spec" json:"spec"`
}
var ListProjectsParamsWrapperDefaultValue = ListProjectsParamsWrapper{
	Kind: "project/ListProjectsParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: project.ListProjectsParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

