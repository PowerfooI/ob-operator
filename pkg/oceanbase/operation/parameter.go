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

package operation

import (
	"fmt"
	"github.com/oceanbase/ob-operator/pkg/oceanbase/const/sql"
	"github.com/oceanbase/ob-operator/pkg/oceanbase/param"
)

func (m *OceanbaseOperationManager) GetParameter(name string, scope *param.Scope) error {
	if scope == nil {
		return m.ExecWithDefaultTimeout(sql.QueryParameter, name)
	} else {
		queryParameterSql := fmt.Sprintf(sql.QueryParameterWithScope, scope.Name)
		return m.ExecWithDefaultTimeout(queryParameterSql, name, scope.Value)
	}
}

func (m *OceanbaseOperationManager) SetParameter(name string, value interface{}, scope *param.Scope) error {
	if scope == nil {
		setParameterSql := fmt.Sprintf(sql.SetParameter, name)
		return m.ExecWithDefaultTimeout(setParameterSql, value)
	} else {
		setParameterSql := fmt.Sprintf(sql.SetParameterWithScope, name, scope.Name)
		return m.ExecWithDefaultTimeout(setParameterSql, value, scope.Value)
	}
}
