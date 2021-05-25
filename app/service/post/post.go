package post

import (
	"NatsunaBlog/app/dao"
	"database/sql"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"math"
)

const (
	PostsInOnePage int = 6
)

//博客简要信息
type GetPostsElement struct {
	Id int			`json:"id"`
	Time gtime.Time	`json:"time"`
	Title string	`json:"title"`
	OnTop bool `json:"ontop"`
	Tag string `json:"tag"`
	Classify string `json:"classify"`
}

//获取总页数
func GetPageNum(r *ghttp.Request) {
	postNum, err := dao.DBBLOGPOST.Count()
	if err != nil{
		r.Response.WritelnExit("GET PAGENUM: " + err.Error())
	}
	r.Response.WritelnExit( int(math.Ceil(float64(postNum) / float64(PostsInOnePage))))
}

//获取单页博客所有内容
func GetOnePost(r *ghttp.Request) {
	postID := r.GetInt("id")
	post, err := dao.DBBLOGPOST.One("id = ?", postID)
	if err != nil {
		r.Response.WritelnExit("GET ONE POST: " + err.Error())
	}
	if post == nil {
		r.Response.WriteExit()
	}
	if r.GetBool("visitOnly") {
		dao.DBBLOGPOST.Update(g.Map{"visit_times": post.VISITTIMES + 1}, "id", postID)
	}
	r.Response.WriteJsonExit(post)
}

//获取当前页博客列表
func GetPosts(r *ghttp.Request) {
	page := r.GetInt("page")
	isAll := r.GetBool("isAll")
	var getCurPosts []*GetPostsElement
	var err error
	if (isAll) {
		err = dao.DBBLOGPOST.
			Order("id desc").
			Limit((page-1)*PostsInOnePage, PostsInOnePage).
			Scan(&getCurPosts)
	} else {
		err = dao.DBBLOGPOST.
			Where("hid = ?", 0).
			Order("ontop desc, id desc").
			Limit((page-1)*PostsInOnePage, PostsInOnePage).
			Scan(&getCurPosts)
	}
	if err != nil {
		r.Response.WritelnExit("GET POSTS: " + err.Error())
	}
	r.Response.WriteJsonExit(getCurPosts)
}

//删除博客
func DeletePost(r *ghttp.Request) {
	if r.Session.GetString("user") == "" {
		r.Response.Writeln("Not Login")
	}
	postID := r.Get("id")
	_, err := dao.DBBLOGPOST.Delete("id = ?", postID)
	if err != nil {
		r.Response.WritelnExit("DELETE: "+ err.Error())
	}
	r.Response.WritelnExit("OK")
}

//更新或新建博客
func UpdatePost(r *ghttp.Request) {
	if r.Session.GetString("user") == "" {
		r.Response.Writeln("Not Login")
	}
	postID := int64(r.GetInt("id"))
	var result sql.Result
	var err error
	title := r.GetString("title")
	if title == "" {
		title = "未命名"
	}
	if postID == -1 {
		result, err = dao.DBBLOGPOST.Insert(g.Map{
			"title": title,
			"content": r.GetString("content"),
			"author": r.Session.GetString("user"),
			"visit_times": 0,
			"hid": r.GetString("hid"),
			"ontop": r.GetString("ontop"),
			"tag": r.GetString("tag"),
			"classify": r.GetString("classify"),
		})
		if err != nil{
			r.Response.WritelnExit("CREATE POSTS: "+ err.Error())
		}
		postID, _ = result.LastInsertId()
	} else {
		result, err = dao.DBBLOGPOST.Update(g.Map{
			"title": title,
			"content": r.GetString("content"),
			"author": r.Session.GetString("user"),
			"hid": r.GetString("hid"),
			"ontop": r.GetString("ontop"),
			"tag": r.GetString("tag"),
			"classify": r.GetString("classify"),
		}, "id", postID)
		if err != nil {
			r.Response.WritelnExit("UPDATE POSTS: "+ err.Error())
		}
	}
	r.Response.WritelnExit(postID)
}
