package admin

import (
	"context"
	"fmt"
	"ipv6Host/internal/dao"
	"ipv6Host/internal/model"
	"ipv6Host/internal/model/entity"
	"ipv6Host/internal/service"
	"ipv6Host/utility"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type (
	sAdmins struct{}
)

func New() *sAdmins {
	return &sAdmins{}
}

func init() {
	service.RegisterAdmins(New())
}

// Create 添加管理员用户
func (s *sAdmins) Create(ctx context.Context, in model.AdminAddInput) (out model.AdminAddOutput, err error) {
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	//加密盐和密码
	UserSalt := grand.S(10)
	in.UserSalt = UserSalt
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	//插入数据，返回id
	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminAddOutput{AdminId: int(lastInsertID)}, err
}

// Delete 删除管理员用户
func (s *sAdmins) Delete(ctx context.Context, id int) (err error) {
	// 删除内容 增加了名称为索引，不能使用逻辑删除，要用物理删除
	_, err = dao.AdminInfo.Ctx(ctx).WherePri(id).Unscoped().Delete()
	return err
}

// GetList 查询管理员列表
// GetList retrieves a list of admin accounts with pagination.
func (s *sAdmins) GetList(ctx context.Context, in model.AdminListInput) (out *model.AdminListOutput, err error) {
	var (
		m = dao.AdminInfo.Ctx(ctx) // Use the admin DAO instead of position
	)

	// Prepare the list model with pagination and optional ordering
	listModel := m.Page(in.Page, in.Size)

	// Optionally: Apply ordering if needed (e.g., by created_at, id, etc.)
	// If no specific order field is requested, we could default to ordering by ID
	listModel = listModel.OrderDesc("id") // Replace "id" with the actual field name for ordering

	var list []*entity.AdminInfo // Use the Admin entity
	if err := listModel.Scan(&list); err != nil {
		return nil, fmt.Errorf("failed to retrieve admin list: %w", err)
	}

	// Count the total number of admin records for pagination
	total, err := m.Count()
	if err != nil {
		return nil, fmt.Errorf("failed to count admins: %w", err)
	}

	// Prepare the output response, including pagination details and the list of admins
	out = &model.AdminListOutput{
		Page:  in.Page,
		Size:  in.Size,
		Total: total,
		List:  make([]model.AdminListItem, len(list)),
	}

	// Convert the list of Admin entities to the list of output items
	for i, item := range list {
		out.List[i] = convertAdminToOutputItem(item)
	}

	return out, nil
}

// convertAdminToOutputItem converts an Admin entity to the corresponding output item
func convertAdminToOutputItem(item *entity.AdminInfo) model.AdminListItem {
	// Adjust the output item fields according to your Admin entity and desired output
	return model.AdminListItem{
		// Populate the fields of the output item with the corresponding fields from the Admin entity
		Id:        item.Id,
		Name:      item.Name,
		IsAdmin:   item.IsAdmin,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
		// ... add other fields as needed
	}
}

// Update 更新管理员用户
func (s *sAdmins) Update(ctx context.Context, in model.AdminUpdateInput) (err error) {

	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	if in.Password != "" {
		UserSalt := grand.S(10)
		in.Password = utility.EncryptPassword(in.Password, UserSalt)
		in.UserSalt = UserSalt
	}
	//更新操作
	_, err = dao.AdminInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.AdminInfo.Columns().Id).
		Where(dao.AdminInfo.Columns().Id, in.Id).
		Update()
	return err

}

// GetAdminByNamePassword 通过用户户名 和 密码查询管理员用户
func (s *sAdmins) GetAdminByNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{} {

	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return nil
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	}
	return g.Map{
		"id":       adminInfo.Id,
		"username": adminInfo.Name,
	}
}
