package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type HostAddReq struct {
	g.Meta      `path:"/backend/host/add" tags:"Host" method:"post" summary:"Add new host"`
	HostID      string `json:"host_id" v:"required#主机唯一标识不能为空" dc:"主机唯一标识"`
	UserID      int    `json:"user_id" v:"required#用户ID不能为空" dc:"用户ID指示该IP地址属于哪个用户"`
	IPv6Address string `json:"ipv6_address" v:"required#主机IPv6地址不能为空" dc:"主机IPv6地址"`
	FreePort    int    `json:"free_port" v:"required#主机空闲端口不能为空" dc:"主机空闲端口"`
}

type HostAddRes struct {
	ID int `json:"id"`
}

type HostDeleteReq struct {
	g.Meta `path:"/backend/host/delete" tags:"Host" method:"post" summary:"Delete host"`
	ID     int `v:"min:1#请选择需要删除的主机" dc:"主机id"`
}

type HostDeleteRes struct{}

type HostGetListCommonReq struct {
	g.Meta `path:"/backend/host/list" tag:"Host List" method:"get" summary:"主机列表"`
	CommonPaginationReq
}

type HostGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type HostUpdateReq struct {
	g.Meta      `path:"/backend/host/update" tags:"Host" method:"post" summary:"Update host"`
	ID          int    `json:"id" v:"min:1#请选择需要修改的主机" dc:"主机id"`
	HostID      string `json:"host_id" v:"required#主机唯一标识不能为空" dc:"主机唯一标识"`
	IPv6Address string `json:"ipv6_address" v:"required#主机IPv6地址不能为空" dc:"主机IPv6地址"`
	FreePort    int    `json:"free_port" v:"required#主机空闲端口不能为空" dc:"主机空闲端口"`
}

type HostUpdateRes struct{}

type Host struct {
	ID          int       `json:"id"`
	HostID      string    `json:"host_id"`
	UserID      int       `json:"user_id"`
	IPv6Address string    `json:"ipv6_address"`
	FreePort    int       `json:"free_port"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}