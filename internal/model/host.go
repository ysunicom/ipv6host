package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type HostCreateUpdateBase struct {
	HostID      string
	IPv6Address string
	FreePort    int
}

type HostCreateInputReq struct {
	HostCreateUpdateBase
}

type HostCreateOutputRes struct {
	ID int `json:"id"`
}

// GfHostGetListInputReq 获取主机列表
type HostGetListInputReq struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GfHostGetListOutputRes 查询列表结果
type HostGetListOutputRes struct {
	List  []HostGetListOutputItem `json:"list" description:"列表"`
	Page  int                       `json:"page" description:"分页码"`
	Size  int                       `json:"size" description:"分页数量"`
	Total int                       `json:"total" description:"数据总数"`
}

// GfHostSearchInputReq 搜索列表
type HostSearchInputReq struct {
	Key  string // 关键字
	Page int    // 分页号码
	Size int    // 分页数量，最大50
	Sort int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// GfHostSearchOutputRes 搜索列表结果
type HostSearchOutputRes struct {
	List  []HostSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int           `json:"stats"` // 搜索统计
	Page  int                      `json:"page"`  // 分页码
	Size  int                      `json:"size"`  // 分页数量
	Total int                      `json:"total"` // 数据总数
}

type HostGetListOutputItem struct {
	Host *HostListItem `json:"host"`
}

type HostSearchOutputItem struct {
	HostGetListOutputItem
}

// GfHostListItem 主要用于列表展示
type HostListItem struct {
	ID          int         `json:"id"`           // 自增ID
	HostID      string      `json:"host_id"`      // 主机唯一标识
	IPv6Address string      `json:"ipv6_address"` // 主机IPv6地址
	FreePort    int         `json:"free_port"`    // 主机空闲端口
	CreatedAt   *gtime.Time `json:"created_at"`   // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"`   // 修改时间
}

type HostUpdateInputReq struct {
	HostCreateUpdateBase
	ID int `json:"id"`
}
