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
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/webhookjob"
	"github.com/spf13/cobra"
	"os"
	"gopkg.in/yaml.v2"
)

var webhookjobListWebhookJobsCmd= &cobra.Command{
	Use:   "listwebhookjobs",
	Short: "listwebhookjobs webhookjob",
	Long:  `Sub-command for ListWebhookJobs`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&ListWebhookJobsParamsWrapperDefaultValue)
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
			
			"PolicyID": "int64",
			
			"ProjectNameOrID": "string",
			
			"Q": "*string",
			
			"Sort": "*string",
			
			"Status": "[]string",
			
			"XIsResourceName": "*bool",
			
			"timeout": "time.Duration",
			
		
		}

		
		localListWebhookJobsParams := &ListWebhookJobsParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localListWebhookJobsParams, err = UnmarshalFileFlag(cmd, localListWebhookJobsParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localListWebhookJobsParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("webhookjob") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*webhookjob.Client).ListWebhookJobs(cliContext, &localListWebhookJobsParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	webhookjobListWebhookJobsCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	webhookjobListWebhookJobsCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	webhookjobListWebhookJobsCmd.Flags().StringP("Page", "", "", "Type: *int64")
		
	webhookjobListWebhookJobsCmd.Flags().StringP("PageSize", "", "", "Type: *int64")
		
	webhookjobListWebhookJobsCmd.Flags().StringP("PolicyID", "", "", "Type: int64")
		
	webhookjobListWebhookJobsCmd.Flags().StringP("ProjectNameOrID", "", "", "Type: string")
		
	webhookjobListWebhookJobsCmd.Flags().StringP("Q", "", "", "Type: *string")
		
	webhookjobListWebhookJobsCmd.Flags().StringP("Sort", "", "", "Type: *string")
		
	webhookjobListWebhookJobsCmd.Flags().StringP("Status", "", "", "Type: []string")
		
	webhookjobListWebhookJobsCmd.Flags().StringP("XIsResourceName", "", "", "Type: *bool")
		
	webhookjobListWebhookJobsCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	webhookjobCmd.AddCommand(webhookjobListWebhookJobsCmd)

}


type ListWebhookJobsParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       webhookjob.ListWebhookJobsParams	`yaml:"spec" json:"spec"`
}
var ListWebhookJobsParamsWrapperDefaultValue = ListWebhookJobsParamsWrapper{
	Kind: "webhookjob/ListWebhookJobsParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: webhookjob.ListWebhookJobsParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

