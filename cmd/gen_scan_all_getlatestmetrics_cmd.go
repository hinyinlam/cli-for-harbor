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
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/scan_all"
	"github.com/spf13/cobra"
	"os"
	"gopkg.in/yaml.v2"
)

var scan_allGetLatestMetricsCmd= &cobra.Command{
	Use:   "getlatestmetrics",
	Short: "getlatestmetrics scan_all",
	Long:  `Sub-command for GetLatestMetrics`,
	Run: func(cmd *cobra.Command, args []string) {

		if isShow, err := cmd.Flags().GetBool("show-request-yaml"); err == nil && isShow == true{
		
			out, err := yaml.Marshal(&GetLatestScanAllMetricsParamsWrapperDefaultValue)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		
			fmt.Println("show-rquest-yaml done")
			os.Exit(0)
		}
		
		listOfCmdParams := map[string]string{
		
			
		
			
			"timeout": "time.Duration",
			
		
		}

		
		localGetLatestScanAllMetricsParams := &GetLatestScanAllMetricsParamsWrapper{}
		filename, err := cmd.Flags().GetString("filename")
		if err == nil && filename!=""{
			localGetLatestScanAllMetricsParams, err = UnmarshalFileFlag(cmd, localGetLatestScanAllMetricsParams)
		}else{
			fitCmdParamIntoDataStruct(cmd, listOfCmdParams, &localGetLatestScanAllMetricsParams.Spec)
		}
		
		
		var apiHandle any = GetApiHandleByModuleName("scan_all") //This is a mapper in runtime to get API client handler, so code generator does a lot less work
        response, err := apiHandle.(*scan_all.Client).GetLatestScanAllMetrics(cliContext, &localGetLatestScanAllMetricsParams.Spec)

		if err != nil {
		    fmt.Println(err)
            os.Exit(-1)
        }

		jsonStr := JsonMarshallWithMultipleTypes(response)
		fmt.Println(jsonStr)

	},
}

func init() {

	scan_allGetLatestMetricsCmd.Flags().StringP("filename", "f", "", "filename to YAML file, all other parameters are ignored if used, use --show-request-yaml for examples")
	var discard bool
	scan_allGetLatestMetricsCmd.Flags().BoolVarP(&discard, "show-request-yaml", "", false, "show example YAML file for --filename param")

	
		
	
		
	scan_allGetLatestMetricsCmd.Flags().StringP("timeout", "", "", "Type: time.Duration")
		
	

	scan_allCmd.AddCommand(scan_allGetLatestMetricsCmd)

}


type GetLatestScanAllMetricsParamsWrapper struct {
	Kind       string                     `yaml:"kind" json:"kind"`
	ApiVersion string                     `yaml:"apiVersion" json:"spec"`
	Metadata   map[string]string          `yaml:"metadata" json:"metadata"`
	Spec       scan_all.GetLatestScanAllMetricsParams	`yaml:"spec" json:"spec"`
}
var GetLatestScanAllMetricsParamsWrapperDefaultValue = GetLatestScanAllMetricsParamsWrapper{
	Kind: "scan_all/GetLatestScanAllMetricsParams",
	ApiVersion: "v1alpha1",
	Metadata: make(map[string]string,0),
	Spec: scan_all.GetLatestScanAllMetricsParams{}, //TODO: Recursively zero all fields so that --show-request-yaml will be able to print some example values
}

