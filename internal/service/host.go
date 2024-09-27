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
	IHosts interface {
		// Create 添加主机
		Create(ctx context.Context, in model.HostCreateInputReq) (out model.HostCreateOutputRes, err error)
		// Delete 删除主机
		Delete(ctx context.Context, id int) (err error)
		// GetList 查询主机列表
		GetList(ctx context.Context, in model.HostGetListInputReq) (out *model.HostGetListOutputRes, err error)
		// Update 更新主机信息
		Update(ctx context.Context, in model.HostUpdateInputReq) (err error)
		Exists(ctx context.Context, in model.HostCreateInputReq) (int, error)
	}
)

var (
	localHosts IHosts
)

func Hosts() IHosts {
	if localHosts == nil {
		panic("implement not found for interface IHosts, forgot register?")
	}
	return localHosts
}

func RegisterHosts(i IHosts) {
	localHosts = i
}
