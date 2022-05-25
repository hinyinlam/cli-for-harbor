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
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/scan"
	"github.com/spf13/cobra"
	"os"
	"gopkg.in/yaml.v2"
)

var scanStopArtifactCmd= &cobra.Command{
	Use:   "stopartifact",
	Short: "stopartifact scan",
	Long:  `Sub-command for StopArtifact`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&StopScanArtifactParamsWrapperDefaultValue)
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
			
			"timeout": "time.Duration",
			
		
		}

		
		localStopScanArtifactParams := &StopScanArtifactParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localStopScanArtifactParams, err = UnmarshalFileFlag(cmd, localStopScanArtifactParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localStopScanArtifactParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("scan") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*scan.Client).StopScanArtifact(cliContext, &localStopScanArtifactParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	scanStopArtifactCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	scanStopArtifactCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	scanStopArtifactCmd.Flags().StringP("ProjectName", "", "", "Type: string")
		
	scanStopArtifactCmd.Flags().StringP("Reference", "", "", "Type: string")
		
	scanStopArtifactCmd.Flags().StringP("RepositoryName", "", "", "Type: string")
		
	scanStopArtifactCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	scanCmd.AddCommand(scanStopArtifactCmd)

}


type StopScanArtifactParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       scan.StopScanArtifactParams	`yaml:"spec" json:"spec"`
}
var StopScanArtifactParamsWrapperDefaultValue = StopScanArtifactParamsWrapper{
	Kind: "scan/StopScanArtifactParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: scan.StopScanArtifactParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

