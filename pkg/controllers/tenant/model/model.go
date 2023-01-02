/*
Copyright (c) 2021 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

package model

type GvTenantList struct {
	TenantID      int64
	TenantName    string
	ZoneList      string
	PrimaryZone   string
	CollationType int64
	ReadOnly      int64
	Locality      string
}

type TenantList struct {
	TenantID          int64
	TenantName        string
	ZoneList          string
	PrimaryZone       string
	ReplicaNum        int64
	LogonlyReplicaNum int64
	Status            string
}

type PoolList struct {
	ResourcePoolID int64
	Name           string
	UnitCount      int64
	UnitConfigID   int64
	ZoneList       string
	TenantID       int64
}

type UnitList struct {
	UnitID             int64
	ResourcePoolID     int64
	Zone               string
	SvrIP              string
	SvrPort            int64
	MigrateFromSvrIP   string
	MigrateFromSvrPort int64
	Status             string
}

type UnitConfigList struct {
	UnitConfigID  int64
	Name          string
	MaxCPU        float64
	MinCPU        float64
	MaxMemory     int64
	MinMemory     int64
	MaxIops       int64
	MinIops       int64
	MaxDiskSize   int64
	MaxSessionNum int64
}

type Resource struct {
	CPUTotal  float64
	MemTotal  int64
	DiskTotal int64
}

type Charset struct {
	Charset string
}

type SysVariableStat struct {
	TenantID int64
	Zone     string
	Name     string
	Value    string
}
