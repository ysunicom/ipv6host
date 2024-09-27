package host

import (
	"context"
	"database/sql"
	"ipv6Host/internal/dao"
	"ipv6Host/internal/model"
	"ipv6Host/internal/model/entity"
	"ipv6Host/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sHosts struct{}
)

func New() *sHosts {
	return &sHosts{}
}

func init() {
	service.RegisterHosts(New())
}

// Create 添加主机
func (s *sHosts) Create(ctx context.Context, in model.HostCreateInputReq) (out model.HostCreateOutputRes, err error) {
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	// 通过用户名查询 UserID
	var admin entity.AdminInfo
	err = dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, in.Name).Scan(&admin)
	if err != nil {
		return out, err
	}
	if admin.Id == 0 {
		return out, gerror.New("用户不存在")
	}

	// 创建主机数据
	hostData := g.Map{
		"host_id":      in.HostID,
		"user_id":      admin.Id,
		"ipv6_address": in.IPv6Address,
		"free_port":    in.FreePort,
	}

	lastInsertID, err := dao.Hosts.Ctx(ctx).Data(hostData).InsertAndGetId()
	return model.HostCreateOutputRes{ID: int(lastInsertID)}, err
}

// Delete 删除主机
func (s *sHosts) Delete(ctx context.Context, id int) (err error) {
	return dao.Hosts.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Hosts.Ctx(ctx).Where(g.Map{
			dao.Hosts.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// GetList 查询主机列表
func (s *sHosts) GetList(ctx context.Context, in model.HostGetListInputReq) (out *model.HostGetListOutputRes, err error) {
	var (
		m = dao.Hosts.Ctx(ctx)
	)
	out = &model.HostGetListOutputRes{
		Page: in.Page,
		Size: in.Size,
	}
	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.Hosts.Columns().Id)
	// 执行查询
	var list []*entity.Hosts
	if err := listModel.Scan(&list); err != nil {
		return nil, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Host
	if err := listModel.ScanList(&out.List, "Host"); err != nil {
		return out, err
	}
	return
}

// Update 更新主机信息
func (s *sHosts) Update(ctx context.Context, in model.HostUpdateInputReq) (err error) {
	return dao.Hosts.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		_, err = dao.Hosts.
			Ctx(ctx).
			OmitEmpty().
			Data(in).
			FieldsEx(dao.Hosts.Columns().Id).
			Where(g.Map{
				dao.Hosts.Columns().Id: in.ID,
			}).Update()
		return err
	})
}

// Exists 检查主机是否存在
func (s *sHosts) Exists(ctx context.Context, in model.HostCreateInputReq) (int, error) {
	var host entity.Hosts
	err := dao.Hosts.Ctx(ctx).Where(dao.Hosts.Columns().HostId, in.HostID).Scan(&host)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // 如果没有记录，返回0和nil
		}
		return 0, err // 其他错误，返回0和错误
	}
	if host.Id > 0 {
		return host.Id, nil
	}
	return 0, nil
}
