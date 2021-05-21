package post

import (
	"NatsunaBlog/app/dao"
	"database/sql"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	_ "github.com/mattn/go-sqlite3"
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

//博客完整信息
/*type PostType struct {
	Time gtime.Time	`json:"time" orm:"time"`
	Title string	`json:"title" orm:"title"`
	Content string	`json:"content" orm:"content"`
	VisitTimes int	`json:"visit_times" orm:"visit_times"`
	Author string	`json:"author" orm:"author"`
	Hided bool `json:"hided" orm:"hided"`
	OnTop bool `json:"ontop" orm:"ontop"`
	Tag string `json:"tag" orm:"tag"`
	Classify string `json:"classify" orm:"classify"`
	Like string `json:"like" orm:"like"`
}*/

/*//博客列表
type GetPostsRes struct{
	Posts []GetPostsElement `json:"posts"`
}*/

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
	postID := r.Get("id")
	post, err := dao.DBBLOGPOST.One("id = ?", postID)
	if err != nil {
		r.Response.WritelnExit("GET ONE POST: " + err.Error())
	}
	if post == nil {
		r.Response.WritelnExit()
	}
	if r.GetBool("visitOnly") {
		dao.DBBLOGPOST.Update(g.Map{"visit_times": post.VisitTimes + 1}, "id", postID)
	}
	r.Response.WriteJsonExit(post)
}

//获取当前页博客列表
func GetPosts(r *ghttp.Request) {
	page := r.GetInt("page")
	var getCurPosts []*GetPostsElement
	err := dao.DBBLOGPOST.Order("ontop desc, id desc").Limit((page-1)*PostsInOnePage, PostsInOnePage).Scan(&getCurPosts)
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
			"hided": r.GetString("hided"),
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
			"hided": r.GetString("hided"),
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
