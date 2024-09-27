// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Hosts is the golang structure for table hosts.
type Hosts struct {
	Id          int         `json:"id"          orm:"id"           description:""`
	HostId      string      `json:"hostId"      orm:"host_id"      description:"主机唯一标识"`
	UserId      int         `json:"userId"      orm:"user_id"      description:""`
	Ipv6Address string      `json:"ipv6Address" orm:"ipv6_address" description:"主机IPv6地址"`
	FreePort    int         `json:"freePort"    orm:"free_port"    description:"主机空闲端口"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:""`
}
