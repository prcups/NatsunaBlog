package post

import (
	"blog/app/dao"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"math"
)

const (
	PostsInOnePage int = 6
)

// 博客简要信息
type GetPostsElement struct {
	Id       int    `json:"id"`
	Time     string `json:"time"`
	Title    string `json:"title"`
	OnTop    bool   `json:"ontop"`
	Tag      string `json:"tag"`
	Classify string `json:"classify"`
}

// 博客完整信息
type GetOnePostElement struct {
	Id         int    `json:"id"`
	Time       string `json:"time"`
	TimeLine   string `json:"timeline"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	OnTop      bool   `json:"ontop"`
	Tag        string `json:"tag"`
	Classify   string `json:"classify"`
	Hid        int    `json:"hid"`
	VisitTimes int    `json:"visit_times"`
	Author     string `json:"author"`
}

// 获取总页数
func GetPageNum(r *ghttp.Request) {
	isAll := r.Get("isAll").Bool()
	var postNum int
	var err error
	if !isAll {
		postNum, err = dao.DbBlogPost.Ctx(r.GetCtx()).
			Where("hid = ?", 0).
			Count()
	} else {
		postNum, err = dao.DbBlogPost.Ctx(r.GetCtx()).Count()
	}
	if err != nil {
		r.Response.WritelnExit("GET PAGENUM: " + err.Error())
	}
	r.Response.WritelnExit(int(math.Ceil(float64(postNum) / float64(PostsInOnePage))))
}

// 获取单页博客所有内容
func GetOnePost(r *ghttp.Request) {
	postID := r.Get("id").Int()
	var post *GetOnePostElement
	err := dao.DbBlogPost.Ctx(r.GetCtx()).Where("id = ?", postID).Scan(&post)
	if err != nil {
		r.Response.WritelnExit("GET ONE POST: " + err.Error())
	}
	if post == nil {
		r.Response.WriteExit()
	}
	if r.Get("visitOnly").Bool() {
		_, _ = dao.DbBlogPost.Ctx(r.GetCtx()).Update(g.Map{"visit_times": post.VisitTimes + 1}, "id", postID)
	}
	r.Response.WriteJsonExit(post)
}

// 获取当前页博客列表
func GetPosts(r *ghttp.Request) {
	page := r.Get("page").Int()
	isAll := r.Get("isAll").Bool()
	var getCurPosts []*GetPostsElement
	var err error
	if isAll {
		err = dao.DbBlogPost.Ctx(r.GetCtx()).
			Order("id desc").
			Limit((page-1)*PostsInOnePage, PostsInOnePage).
			Scan(&getCurPosts)
	} else {
		err = dao.DbBlogPost.Ctx(r.GetCtx()).
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

// 删除博客
func DeletePost(r *ghttp.Request) {
	if s, _ := r.Session.Get("user"); s.String() == "" {
		r.Response.Writeln("Not Login")
	}
	postID := r.Get("id")
	_, err := dao.DbBlogPost.Ctx(r.GetCtx()).Delete("id = ?", postID)
	if err != nil {
		r.Response.WritelnExit("DELETE: " + err.Error())
	}
	r.Response.WritelnExit("OK")
}

// 更新或新建博客
func UpdatePost(r *ghttp.Request) {
	if s, _ := r.Session.Get("user"); s.String() == "" {
		r.Response.Writeln("Not Login")
	}
	postID := int64(r.Get("id").Int())
	var result sql.Result
	var err error
	title := r.Get("title").String()
	if title == "" {
		title = "未命名"
	}
	hid := r.Get("hid").Int()
	if hid != 0 && hid != 1 {
		hid = 0
	}
	ontop := r.Get("ontop").Int()
	if ontop != 0 && ontop != 1 {
		ontop = 0
	}
	classify := r.Get("classify").String()
	if classify == "" {
		classify = "默认分类"
	}

	if postID == -1 {
		authorVar, _ := r.Session.Get("user")
		result, err = dao.DbBlogPost.Ctx(r.GetCtx()).Insert(g.Map{
			"title":       title,
			"content":     r.Get("content").String(),
			"author":      authorVar.String(),
			"visit_times": 0,
			"hid":         hid,
			"ontop":       ontop,
			"tag":         r.Get("tag").String(),
			"classify":    classify,
		})
		if err != nil {
			r.Response.WritelnExit("CREATE POSTS: " + err.Error())
		}
		postID, _ = result.LastInsertId()
	} else {
		authorVar, _ := r.Session.Get("user")
		result, err = dao.DbBlogPost.Ctx(r.GetCtx()).Update(g.Map{
			"title":    title,
			"content":  r.Get("content").String(),
			"author":   authorVar.String(),
			"hid":      hid,
			"ontop":    ontop,
			"tag":      r.Get("tag").String(),
			"classify": classify,
		}, "id", postID)
		if err != nil {
			r.Response.WritelnExit("UPDATE POSTS: " + err.Error())
		}
	}
	r.Response.WritelnExit(postID)
}
