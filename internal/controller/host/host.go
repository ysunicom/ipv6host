package host

import (
	"context"
	"ipv6Host/api/backend"
	"ipv6Host/internal/model"
	"ipv6Host/internal/service"
)

// 主机管理
var Host = cHost{}

type cHost struct{}

func New() *cHost {
	return &cHost{}
}

func (c *cHost) Create(ctx context.Context, req *backend.HostAddReq) (res *backend.HostAddRes, err error) {
	// 检查HostID是否存在
	exists, err := service.Hosts().Exists(ctx, model.HostCreateInputReq{
		HostCreateUpdateBase: model.HostCreateUpdateBase{
			HostID:      req.HostID,
			IPv6Address: req.IPv6Address,
			Name:        req.Name,
			FreePort:    req.FreePort,
		},
	})
	if err != nil {
		return nil, err
	}

	if exists != 0 {
		// 如果存在，更新记录
		err = service.Hosts().Update(ctx, model.HostUpdateInputReq{
			ID: exists,
			HostCreateUpdateBase: model.HostCreateUpdateBase{
				HostID:      req.HostID,
				IPv6Address: req.IPv6Address,
				Name:        req.Name,
				FreePort:    req.FreePort,
			},
		})
		if err != nil {
			return nil, err
		}
		return &backend.HostAddRes{ID: exists}, nil
	}

	// 如果不存在，创建新记录
	out, err := service.Hosts().Create(ctx, model.HostCreateInputReq{
		HostCreateUpdateBase: model.HostCreateUpdateBase{
			HostID:      req.HostID,
			IPv6Address: req.IPv6Address,
			Name:        req.Name,
			FreePort:    req.FreePort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.HostAddRes{ID: out.ID}, nil
}

// 删除主机
func (c *cHost) Delete(ctx context.Context, req *backend.HostDeleteReq) (res *backend.HostDeleteRes, err error) {
	err = service.Hosts().Delete(ctx, req.ID)
	return
}

// 更新主机
func (c *cHost) Update(ctx context.Context, req *backend.HostUpdateReq) (res *backend.HostUpdateRes, err error) {
	err = service.Hosts().Update(ctx, model.HostUpdateInputReq{
		ID: req.ID,
		HostCreateUpdateBase: model.HostCreateUpdateBase{
			HostID:      req.HostID,
			Name:        req.Name,
			IPv6Address: req.IPv6Address,
			FreePort:    req.FreePort,
		},
	})
	return
}

// List 主机列表
func (c *cHost) List(ctx context.Context, req *backend.HostGetListCommonReq) (res *backend.HostGetListCommonRes, err error) {
	getListRes, err := service.Hosts().GetList(ctx, model.HostGetListInputReq{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.HostGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
