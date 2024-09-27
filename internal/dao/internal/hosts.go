// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HostsDao is the data access object for table gf_hosts.
type HostsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns HostsColumns // columns contains all the column names of Table for convenient usage.
}

// HostsColumns defines and stores column names for table gf_hosts.
type HostsColumns struct {
	Id          string //
	HostId      string // 主机唯一标识
	Ipv6Address string // 主机IPv6地址
	FreePort    string // 主机空闲端口
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
}

// hostsColumns holds the columns for table gf_hosts.
var hostsColumns = HostsColumns{
	Id:          "id",
	HostId:      "host_id",
	Ipv6Address: "ipv6_address",
	FreePort:    "free_port",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewHostsDao creates and returns a new DAO object for table data access.
func NewHostsDao() *HostsDao {
	return &HostsDao{
		group:   "default",
		table:   "gf_hosts",
		columns: hostsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HostsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HostsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HostsDao) Columns() HostsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HostsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HostsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HostsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
