// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// DBBLOGCOMMENT is the golang structure for table DB_BLOG_COMMENT.
type DBBLOGCOMMENT struct {
	Id       int         `orm:"id"       json:"id"`       //
	Blogid   int         `orm:"blogid"   json:"blogid"`   //
	Username string      `orm:"username" json:"username"` //
	Content  string      `orm:"content"  json:"content"`  //
	Email    string      `orm:"email"    json:"email"`    //
	Time     *gtime.Time `orm:"time"     json:"time"`     //
	Like     int         `orm:"like"     json:"like"`     //
}
