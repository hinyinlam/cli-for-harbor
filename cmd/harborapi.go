/*
 * CLI for Harbor
 * Copyright 2022 VMware, Inc.
 *
 * This product is licensed to you under the Apache 2.0 license (the "License").  You may not use this product except in compliance with the Apache 2.0 License.
 *
 * This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
 */

/*
 * Author: Hin Lam <hinl@vmware.com>
 */

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var clientSetConfig = harbor.ClientSetConfig{}

var harborApi *client.HarborAPI
var cliContext context.Context

func init() {
	cliContext = context.Background()
}

var AuthFile string

func readAuthInfo() {
	authData, err := ioutil.ReadFile(AuthFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}

	var clientSetConfigWrapper = &ClientSetConfigWrapper{}
	err = yaml.Unmarshal(authData, &clientSetConfigWrapper)
	if err != nil {
		panic(err)
	}
	clientSetConfig = clientSetConfigWrapper.Spec
}

type ClientSetConfigWrapper struct {
	Kind       string
	ApiVersion string
	Spec       harbor.ClientSetConfig
}

func SaveAuthInfo(csc harbor.ClientSetConfig) {
	cscWrapper := ClientSetConfigWrapper{
		Kind:       "ClientSetConfigParam",
		ApiVersion: "v2alpha1",
		Spec:       csc,
	}
	bytes, err := yaml.Marshal(cscWrapper)
	if err != nil {
		panic(err)
	}
	os.MkdirAll(filepath.Dir(AuthFile), os.ModePerm)
	err = ioutil.WriteFile(AuthFile, bytes, 0600)
	if err != nil {
		panic(err)
	}
	fmt.Println("Auth info saved to ", AuthFile)
}

func InitHarborApi() {
	readAuthInfo()
	cs, err := harbor.NewClientSet(&clientSetConfig)
	if err != nil {
		panic(err)
	}
	harborApi = cs.V2()
}

func JsonMarshallWithMultipleTypes(v any) string {
	if v == nil {
		return ""
	}
	m := make(map[string]interface{})
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	v, ok := m["Payload"]
	if ok {
		if output, err := json.MarshalIndent(v, "", "    "); err == nil {
			return string(output)
		}
	}

	myerror, ok := v.(error)
	if ok {
		str := myerror.Error()
		if output, err := json.MarshalIndent(str, "", "    "); err == nil {
			return string(output)
		}
	}
	return fmt.Sprintf("%v", v)
	// panic("Not sure how to deal with response message in JSON Marshall for this API call")
	// return ""
}

//This is a simple and maunual to resolve which harborApi.APIName we should use
//Use of reflection could be better but significantly complex
func GetApiHandleByModuleName(moduleName string) any {
	switch moduleName {
	case "artifact":
		return harborApi.Artifact
	case "auditlog":
		return harborApi.Auditlog
	case "configure":
		return harborApi.Configure
	case "gc":
		return harborApi.GC
	case "health":
		return harborApi.Health
	case "icon":
		return harborApi.Icon
	case "immutable":
		return harborApi.Immutable
	case "label":
		return harborApi.Label
	case "ldap":
		return harborApi.Ldap
	case "member":
		return harborApi.Member
	case "oidc":
		return harborApi.OIDC
	case "ping":
		return harborApi.Ping
	case "preheat":
		return harborApi.Preheat
	case "project":
		return harborApi.Project
	case "project_metadata":
		return harborApi.ProjectMetadata
	case "quota":
		return harborApi.Quota
	case "registry":
		return harborApi.Registry
	case "replication":
		return harborApi.Replication
	case "repository":
		return harborApi.Repository
	case "retention":
		return harborApi.Retention
	case "robot":
		return harborApi.Robot
	case "robotv1":
		return harborApi.Robotv1
	case "scan":
		return harborApi.Scan
	case "scan_all":
		return harborApi.ScanAll
	case "scanner":
		return harborApi.Scanner
	case "search":
		return harborApi.Search
	case "statistic":
		return harborApi.Statistic
	case "system_cve_allowlist":
		return harborApi.SystemCVEAllowlist
	case "systeminfo":
		return harborApi.Systeminfo
	case "user":
		return harborApi.User
	case "usergroup":
		return harborApi.Usergroup
	case "webhook":
		return harborApi.Webhook
	case "webhookjob":
		return harborApi.Webhookjob
	case "runtime.ClientTransport":
		return harborApi.Transport
	default:
		log.Panicf("Failed to find mapping between moduleName: %v and APIClient Struct", moduleName)
		panic("Mapping should be updated in harborapi.go file")
	}

}

func fitCmdParamIntoDataStruct[T any](cmd *cobra.Command, listOfArgs map[string]string, object *T) {
	paramMap := make(map[string]any, 0)
	for arg, argType := range listOfArgs {
		argValue, err := cmd.Flags().GetString(arg)
		if err == nil && argValue != "" {
			switch argType {
			case "*int64":
				tmp, err := strconv.ParseInt(argValue, 10, 64)
				if err != nil {
					panic(err)
				}
				paramMap[arg] = tmp
			default:
				paramMap[arg] = argValue
			}
		}
	}
	out, err := json.Marshal(paramMap)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(out, object)
	if err != nil {
		panic(err)
	}
}

func UnmarshalFileFlag[T any](cmd *cobra.Command, object *T) (*T, error) {
	filename, err := cmd.Flags().GetString("filename")
	if err != nil {
		return object, nil
	}
	if filename == "" {
		return object, nil
	}

	paramData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(paramData, &object)
	if err != nil {
		panic(err)
	}
	return object, nil
}
