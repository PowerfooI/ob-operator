/*
Copyright (c) 2023 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

package v2alpha1

type Parameter struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	ScopeType string `json:"scopeType,omitempty"`
	Scope     string `json:"scope,omitempty"`
}

type ParameterValue struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Zone   string `json:"zone"`
	Server string `json:"server"`
}