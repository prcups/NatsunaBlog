// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DbBlogConfigDao is the data access object for the table DB_BLOG_CONFIG.
type DbBlogConfigDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  DbBlogConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// DbBlogConfigColumns defines and stores column names for the table DB_BLOG_CONFIG.
type DbBlogConfigColumns struct {
	Username string //
	Password string //
}

// dbBlogConfigColumns holds the columns for the table DB_BLOG_CONFIG.
var dbBlogConfigColumns = DbBlogConfigColumns{
	Username: "USERNAME",
	Password: "PASSWORD",
}

// NewDbBlogConfigDao creates and returns a new DAO object for table data access.
func NewDbBlogConfigDao(handlers ...gdb.ModelHandler) *DbBlogConfigDao {
	return &DbBlogConfigDao{
		group:    "default",
		table:    "DB_BLOG_CONFIG",
		columns:  dbBlogConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DbBlogConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DbBlogConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DbBlogConfigDao) Columns() DbBlogConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DbBlogConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DbBlogConfigDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DbBlogConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
