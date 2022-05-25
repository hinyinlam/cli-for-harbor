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

var artifactGetCmd= &cobra.Command{
	Use:   "get",
	Short: "get artifact",
	Long:  `Sub-command for Get`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&GetArtifactParamsWrapperDefaultValue)
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
			
			"Reference": "string",
			
			"RepositoryName": "string",
			
			"WithAccessory": "*bool",
			
			"WithImmutableStatus": "*bool",
			
			"WithLabel": "*bool",
			
			"WithScanOverview": "*bool",
			
			"WithSignature": "*bool",
			
			"WithTag": "*bool",
			
			"XAcceptVulnerabilities": "*string",
			
			"timeout": "time.Duration",
			
		
		}

		
		localGetArtifactParams := &GetArtifactParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localGetArtifactParams, err = UnmarshalFileFlag(cmd, localGetArtifactParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localGetArtifactParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("artifact") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*artifact.Client).GetArtifact(cliContext, &localGetArtifactParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	artifactGetCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	artifactGetCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	artifactGetCmd.Flags().StringP("Page", "", "", "Type: *int64")
		
	artifactGetCmd.Flags().StringP("PageSize", "", "", "Type: *int64")
		
	artifactGetCmd.Flags().StringP("ProjectName", "", "", "Type: string")
		
	artifactGetCmd.Flags().StringP("Reference", "", "", "Type: string")
		
	artifactGetCmd.Flags().StringP("RepositoryName", "", "", "Type: string")
		
	artifactGetCmd.Flags().StringP("WithAccessory", "", "", "Type: *bool")
		
	artifactGetCmd.Flags().StringP("WithImmutableStatus", "", "", "Type: *bool")
		
	artifactGetCmd.Flags().StringP("WithLabel", "", "", "Type: *bool")
		
	artifactGetCmd.Flags().StringP("WithScanOverview", "", "", "Type: *bool")
		
	artifactGetCmd.Flags().StringP("WithSignature", "", "", "Type: *bool")
		
	artifactGetCmd.Flags().StringP("WithTag", "", "", "Type: *bool")
		
	artifactGetCmd.Flags().StringP("XAcceptVulnerabilities", "", "", "Type: *string")
		
	artifactGetCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	artifactCmd.AddCommand(artifactGetCmd)

}


type GetArtifactParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       artifact.GetArtifactParams	`yaml:"spec" json:"spec"`
}
var GetArtifactParamsWrapperDefaultValue = GetArtifactParamsWrapper{
	Kind: "artifact/GetArtifactParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: artifact.GetArtifactParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

