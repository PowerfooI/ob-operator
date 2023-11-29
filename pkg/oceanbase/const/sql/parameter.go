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

package sql

const (
	SetParameter            = "alter system set %s = ?"
	SetParameterWithScope   = "alter system set %s = ? %s = ?"
	QueryParameter          = "select zone, svr_ip, svr_port, name, value, scope, edit_level from __all_virtual_sys_parameter_stat where name = ?"
	QueryParameterWithScope = "select zone, svr_ip, svr_port, name, value, scope, edit_level from __all_virtual_sys_parameter_stat where name = ? and %s = ?"
)

const (
	ListParametersWithTenantID = "select name, value from GV$OB_PARAMETERS where tenant_id = ?"
	SelectCompatibleOfTenants  = "select name, value, tenant_id from GV$OB_PARAMETERS where name = 'compatible'"
)
