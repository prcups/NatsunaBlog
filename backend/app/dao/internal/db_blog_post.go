// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DbBlogPostDao is the data access object for the table DB_BLOG_POST.
type DbBlogPostDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DbBlogPostColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DbBlogPostColumns defines and stores column names for the table DB_BLOG_POST.
type DbBlogPostColumns struct {
	Id         string //
	Title      string //
	Timeline   string //
	Time       string //
	Author     string //
	Hid        string //
	Content    string //
	Classify   string //
	Tag        string //
	VisitTimes string //
	Ontop      string //
}

// dbBlogPostColumns holds the columns for the table DB_BLOG_POST.
var dbBlogPostColumns = DbBlogPostColumns{
	Id:         "ID",
	Title:      "TITLE",
	Timeline:   "TIMELINE",
	Time:       "TIME",
	Author:     "AUTHOR",
	Hid:        "HID",
	Content:    "CONTENT",
	Classify:   "CLASSIFY",
	Tag:        "TAG",
	VisitTimes: "VISIT_TIMES",
	Ontop:      "ONTOP",
}

// NewDbBlogPostDao creates and returns a new DAO object for table data access.
func NewDbBlogPostDao(handlers ...gdb.ModelHandler) *DbBlogPostDao {
	return &DbBlogPostDao{
		group:    "default",
		table:    "DB_BLOG_POST",
		columns:  dbBlogPostColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DbBlogPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DbBlogPostDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DbBlogPostDao) Columns() DbBlogPostColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DbBlogPostDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DbBlogPostDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DbBlogPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
