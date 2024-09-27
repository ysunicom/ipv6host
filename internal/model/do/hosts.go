// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Hosts is the golang structure of table gf_hosts for DAO operations like Where/Data.
type Hosts struct {
	g.Meta      `orm:"table:gf_hosts, do:true"`
	Id          interface{} //
	HostId      interface{} // 主机唯一标识
	Ipv6Address interface{} // 主机IPv6地址
	FreePort    interface{} // 主机空闲端口
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
}
