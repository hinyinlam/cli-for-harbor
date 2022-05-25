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

var artifactListCmd= &cobra.Command{
	Use:   "list",
	Short: "list artifact",
	Long:  `Sub-command for List`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&ListArtifactsParamsWrapperDefaultValue)
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
			
			"RepositoryName": "string",
			
			"Sort": "*string",
			
			"WithAccessory": "*bool",
			
			"WithImmutableStatus": "*bool",
			
			"WithLabel": "*bool",
			
			"WithScanOverview": "*bool",
			
			"WithSignature": "*bool",
			
			"WithTag": "*bool",
			
			"XAcceptVulnerabilities": "*string",
			
			"timeout": "time.Duration",
			
		
		}

		
		localListArtifactsParams := &ListArtifactsParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localListArtifactsParams, err = UnmarshalFileFlag(cmd, localListArtifactsParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localListArtifactsParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("artifact") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*artifact.Client).ListArtifacts(cliContext, &localListArtifactsParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	artifactListCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	artifactListCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	artifactListCmd.Flags().StringP("Page", "", "", "Type: *int64")
		
	artifactListCmd.Flags().StringP("PageSize", "", "", "Type: *int64")
		
	artifactListCmd.Flags().StringP("ProjectName", "", "", "Type: string")
		
	artifactListCmd.Flags().StringP("Q", "", "", "Type: *string")
		
	artifactListCmd.Flags().StringP("RepositoryName", "", "", "Type: string")
		
	artifactListCmd.Flags().StringP("Sort", "", "", "Type: *string")
		
	artifactListCmd.Flags().StringP("WithAccessory", "", "", "Type: *bool")
		
	artifactListCmd.Flags().StringP("WithImmutableStatus", "", "", "Type: *bool")
		
	artifactListCmd.Flags().StringP("WithLabel", "", "", "Type: *bool")
		
	artifactListCmd.Flags().StringP("WithScanOverview", "", "", "Type: *bool")
		
	artifactListCmd.Flags().StringP("WithSignature", "", "", "Type: *bool")
		
	artifactListCmd.Flags().StringP("WithTag", "", "", "Type: *bool")
		
	artifactListCmd.Flags().StringP("XAcceptVulnerabilities", "", "", "Type: *string")
		
	artifactListCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	artifactCmd.AddCommand(artifactListCmd)

}


type ListArtifactsParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       artifact.ListArtifactsParams	`yaml:"spec" json:"spec"`
}
var ListArtifactsParamsWrapperDefaultValue = ListArtifactsParamsWrapper{
	Kind: "artifact/ListArtifactsParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: artifact.ListArtifactsParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

