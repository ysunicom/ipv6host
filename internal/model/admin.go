package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminBase 管理员基本信息结构体，用于创建和更新时共用的字段
type AdminBase struct {
	Name     string
	Password string
	RoleIds  int
	UserSalt string
	IsAdmin  int
}

// AdminAddInput 管理员添加请求输入参数
type AdminAddInput struct {
	AdminBase
}

// AdminAddOutput 管理员添加响应输出参数
type AdminAddOutput struct {
	AdminId int `json:"admin_id"`
}

// AdminUpdateInput 管理员更新请求输入参数
type AdminUpdateInput struct {
	Id int
	AdminBase
}

// AdminDeleteInput 管理员删除请求输入参数
type AdminDeleteInput struct {
	Id int
}

// AdminListInput 管理员列表请求输入参数
type AdminListInput struct {
	Page int
	Size int
}

// AdminListOutput 管理员列表响应输出参数
type AdminListOutput struct {
	List  []AdminListItem `json:"list"`
	Page  int             `json:"page"`
	Size  int             `json:"size"`
	Total int             `json:"total"`
}

// AdminListItem 管理员列表项
type AdminListItem struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	RoleIds   int         `json:"role_ids"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
	IsAdmin   int         `json:"is_admin"`
}
