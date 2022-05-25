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
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/robotv1"
	"github.com/spf13/cobra"
	"os"
	"gopkg.in/yaml.v2"
)

var robotv1CreateRobotV1Cmd= &cobra.Command{
	Use:   "createrobotv1",
	Short: "createrobotv1 robotv1",
	Long:  `Sub-command for CreateRobotV1`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&CreateRobotV1ParamsWrapperDefaultValue)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		
			fmt.Println("show-rquest-yaml done")
			os.Exit(0)
		}
		
		listOfCmdParams := map[string]string{
		
			
		
			
			"ProjectNameOrID": "string",
			
			"Robot": "*models.RobotCreateV1",
			
			"XIsResourceName": "*bool",
			
			"timeout": "time.Duration",
			
		
		}

		
		localCreateRobotV1Params := &CreateRobotV1ParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localCreateRobotV1Params, err = UnmarshalFileFlag(cmd, localCreateRobotV1Params)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localCreateRobotV1Params.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("robotv1") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*robotv1.Client).CreateRobotV1(cliContext, &localCreateRobotV1Params.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	robotv1CreateRobotV1Cmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	robotv1CreateRobotV1Cmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	robotv1CreateRobotV1Cmd.Flags().StringP("ProjectNameOrID", "", "", "Type: string")
		
	robotv1CreateRobotV1Cmd.Flags().StringP("Robot", "", "", "Type: *models.RobotCreateV1")
		
	robotv1CreateRobotV1Cmd.Flags().StringP("XIsResourceName", "", "", "Type: *bool")
		
	robotv1CreateRobotV1Cmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	robotv1Cmd.AddCommand(robotv1CreateRobotV1Cmd)

}


type CreateRobotV1ParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       robotv1.CreateRobotV1Params	`yaml:"spec" json:"spec"`
}
var CreateRobotV1ParamsWrapperDefaultValue = CreateRobotV1ParamsWrapper{
	Kind: "robotv1/CreateRobotV1Params",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: robotv1.CreateRobotV1Params{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

