package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"ipv6Host/api/backend"
	"ipv6Host/internal/consts"
	"ipv6Host/internal/model"
	"ipv6Host/internal/service"
)

// Admin 管理员账户管理
var Admin = cAdmin{}

type cAdmin struct{}

func New() *cAdmin {
	return &cAdmin{}
}

// Create 添加管理员账户
func (c *cAdmin) Create(ctx context.Context, req *backend.AdminAddReq) (res *backend.AdminAddRes, err error) {
	out, err := service.Admins().Create(ctx, model.AdminAddInput{
		AdminBase: model.AdminBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminAddRes{AdminId: out.AdminId}, err
}

// Delete 删除管理员账户
func (c *cAdmin) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admins().Delete(ctx, req.Id)
	return
}

// List 查询管理员账户列表
func (c *cAdmin) List(ctx context.Context, req *backend.AdminGetListReq) (res *backend.AdminGetListRes, err error) {
	getListRes, err := service.Admins().GetList(ctx, model.AdminListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	// Note: You might need to convert getListRes.List if the service layer returns a different data structure
	return &backend.AdminGetListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// Update 更新管理员账户信息
func (c *cAdmin) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	err = service.Admins().Update(ctx, model.AdminUpdateInput{
		Id: req.Id,
		AdminBase: model.AdminBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	return
}

// for gtoken
func (c *cAdmin) Info(ctx context.Context, req *backend.AdminGetInfoReq) (res *backend.AdminGetInfoRes, err error) {
	return &backend.AdminGetInfoRes{
		Id:      gconv.Int(ctx.Value(consts.CtxAdminId)),
		Name:    gconv.String(ctx.Value(consts.CtxAdminName)),
		RoleIds: gconv.Int(ctx.Value(consts.CtxAdminRoleIds)),
		IsAdmin: gconv.Int(ctx.Value(consts.CtxAdminIsAdmin)),
	}, nil
}
