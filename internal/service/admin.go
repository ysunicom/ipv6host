// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"ipv6Host/internal/model"
)

type (
	IAdmins interface {
		// Create 添加管理员用户
		Create(ctx context.Context, in model.AdminAddInput) (out model.AdminAddOutput, err error)
		// Delete 删除管理员用户
		Delete(ctx context.Context, id int) (err error)
		// GetList 查询管理员列表
		// GetList retrieves a list of admin accounts with pagination.
		GetList(ctx context.Context, in model.AdminListInput) (out *model.AdminListOutput, err error)
		// Update 更新管理员用户
		Update(ctx context.Context, in model.AdminUpdateInput) (err error)
		// GetAdminByNamePassword 通过用户户名 和 密码查询管理员用户
		GetAdminByNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{}
	}
)

var (
	localAdmins IAdmins
)

func Admins() IAdmins {
	if localAdmins == nil {
		panic("implement not found for interface IAdmins, forgot register?")
	}
	return localAdmins
}

func RegisterAdmins(i IAdmins) {
	localAdmins = i
}
