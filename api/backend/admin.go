package backend

import "github.com/gogf/gf/v2/frame/g"

// AdminAddReq 用于添加管理员的请求结构体
type AdminAddReq struct {
	g.Meta   `path:"/backend/admin/add" tags:"Admin" method:"post" summary:"增加管理员接口"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	RoleIds  int    `json:"role_ids" dc:"角色ids"`
	IsAdmin  int    `json:"is_admin" dc:"是否超级管理员"`
}

// AdminAddRes 用于添加管理员的响应结构体
type AdminAddRes struct {
	AdminId int `json:"admin_id"`
}

// AdminDeleteReq 用于删除管理员的请求结构体
type AdminDeleteReq struct {
	g.Meta `path:"/backend/admin/delete" tags:"Admin" method:"delete" summary:"删除管理员接口"`
	Id     int `v:"min:1#请选择需要删除的管理员" dc:"管理员ID"`
}

// AdminDeleteRes 用于删除管理员的响应结构体
type AdminDeleteRes struct{}

// AdminUpdateReq 用于修改管理员信息的请求结构体
type AdminUpdateReq struct {
	g.Meta   `path:"/backend/admin/update/{Id}" method:"post" tags:"管理员" summary:"修改管理员信息接口"`
	Id       int    `json:"id" v:"min:1#请选择需要修改的管理员" dc:"管理员ID"`
	Name     string `json:"name" dc:"用户名"`
	Password string `json:"password" dc:"密码"`
	RoleIds  int    `json:"role_ids" dc:"角色ids"`
	IsAdmin  int    `json:"is_admin" dc:"是否超级管理员"`
}

// AdminUpdateRes 用于修改管理员信息的响应结构体
type AdminUpdateRes struct{}

// AdminGetListReq 用于获取管理员列表的请求结构体
type AdminGetListReq struct {
	g.Meta `path:"/backend/admin/list" tags:"Admin List" method:"get" summary:"管理员列表"`
	CommonPaginationReq
}

// AdminGetListRes 用于获取管理员列表的响应结构体
type AdminGetListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type AdminGetInfoReq struct {
	g.Meta `path:"/backend/admin/info" method:"get"`
}

// for gtoken
type AdminGetInfoRes struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	RoleIds int    `json:"role_ids"`
	IsAdmin int    `json:"is_admin"`
}

type AdminGetInfoGtokenRes struct {
	Id          int    `json:"id"`
	IdentityKey string `json:"identity_key"`
	Payload     string `json:"payload"`
}
