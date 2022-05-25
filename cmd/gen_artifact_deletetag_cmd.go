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

var artifactDeleteTagCmd= &cobra.Command{
	Use:   "deletetag",
	Short: "deletetag artifact",
	Long:  `Sub-command for DeleteTag`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&DeleteTagParamsWrapperDefaultValue)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		
			fmt.Println("show-rquest-yaml done")
			os.Exit(0)
		}
		
		listOfCmdParams := map[string]string{
		
			
		
			
			"ProjectName": "string",
			
			"Reference": "string",
			
			"RepositoryName": "string",
			
			"TagName": "string",
			
			"timeout": "time.Duration",
			
		
		}

		
		localDeleteTagParams := &DeleteTagParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localDeleteTagParams, err = UnmarshalFileFlag(cmd, localDeleteTagParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localDeleteTagParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("artifact") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*artifact.Client).DeleteTag(cliContext, &localDeleteTagParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	artifactDeleteTagCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	artifactDeleteTagCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	artifactDeleteTagCmd.Flags().StringP("ProjectName", "", "", "Type: string")
		
	artifactDeleteTagCmd.Flags().StringP("Reference", "", "", "Type: string")
		
	artifactDeleteTagCmd.Flags().StringP("RepositoryName", "", "", "Type: string")
		
	artifactDeleteTagCmd.Flags().StringP("TagName", "", "", "Type: string")
		
	artifactDeleteTagCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	artifactCmd.AddCommand(artifactDeleteTagCmd)

}


type DeleteTagParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       artifact.DeleteTagParams	`yaml:"spec" json:"spec"`
}
var DeleteTagParamsWrapperDefaultValue = DeleteTagParamsWrapper{
	Kind: "artifact/DeleteTagParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: artifact.DeleteTagParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

