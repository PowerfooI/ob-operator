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

package task

import (
	cloudv2alpha1 "github.com/oceanbase/ob-operator/api/v2alpha1"
	serverstatus "github.com/oceanbase/ob-operator/pkg/const/status/observer"
	flowname "github.com/oceanbase/ob-operator/pkg/task/const/flow/name"
	taskname "github.com/oceanbase/ob-operator/pkg/task/const/task/name"
)

func CreateServerForBootstrapTaskFlow() *TaskFlow {
	return &TaskFlow{
		OperationContext: &cloudv2alpha1.OperationContext{
			Name:         flowname.CreateServerForBootstrap,
			Tasks:        []string{taskname.CreateOBPVC, taskname.CreateOBPod, taskname.WaitOBPodReady},
			TargetStatus: serverstatus.BootstrapReady,
		},
	}
}

func CreateServerTaskFlow() *TaskFlow {
	return nil
}