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
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/member"
	"github.com/spf13/cobra"
	"os"
	"gopkg.in/yaml.v2"
)

var memberCreateProjectCmd= &cobra.Command{
	Use:   "createproject",
	Short: "createproject member",
	Long:  `Sub-command for CreateProject`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&CreateProjectMemberParamsWrapperDefaultValue)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		
			fmt.Println("show-rquest-yaml done")
			os.Exit(0)
		}
		
		listOfCmdParams := map[string]string{
		
			
		
			
			"ProjectMember": "*models.ProjectMember",
			
			"ProjectNameOrID": "string",
			
			"XIsResourceName": "*bool",
			
			"timeout": "time.Duration",
			
		
		}

		
		localCreateProjectMemberParams := &CreateProjectMemberParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localCreateProjectMemberParams, err = UnmarshalFileFlag(cmd, localCreateProjectMemberParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localCreateProjectMemberParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("member") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*member.Client).CreateProjectMember(cliContext, &localCreateProjectMemberParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	memberCreateProjectCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	memberCreateProjectCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	memberCreateProjectCmd.Flags().StringP("ProjectMember", "", "", "Type: *models.ProjectMember")
		
	memberCreateProjectCmd.Flags().StringP("ProjectNameOrID", "", "", "Type: string")
		
	memberCreateProjectCmd.Flags().StringP("XIsResourceName", "", "", "Type: *bool")
		
	memberCreateProjectCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	memberCmd.AddCommand(memberCreateProjectCmd)

}


type CreateProjectMemberParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       member.CreateProjectMemberParams	`yaml:"spec" json:"spec"`
}
var CreateProjectMemberParamsWrapperDefaultValue = CreateProjectMemberParamsWrapper{
	Kind: "member/CreateProjectMemberParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: member.CreateProjectMemberParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

