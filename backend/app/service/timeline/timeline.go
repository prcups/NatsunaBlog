package timeline

import (
	"blog/app/dao"
	"github.com/gogf/gf/v2/net/ghttp"
)

type GetPostsElement struct {
	Id       int    `json:"id"`
	Time     string `json:"time"`
	Title    string `json:"title"`
	OnTop    bool   `json:"ontop"`
	Tag      string `json:"tag"`
	Classify string `json:"classify"`
}

// 获取所有归档
func GetAllTimeLine(r *ghttp.Request) {
	timeline, err := dao.DbBlogPost.Ctx(r.GetCtx()).Where("hid = ?", 0).Array("distinct timeline")
	if err != nil {
		r.Response.WritelnExit("GETALLTIMELINE: " + err.Error())
	}
	r.Response.WriteJsonExit(timeline)
}

// 获取单个归档下的内容
func GetPostsOfTimeLine(r *ghttp.Request) {
	timeline := r.Get("timeline").String()
	var posts []GetPostsElement
	dao.DbBlogPost.Ctx(r.GetCtx()).Where("hid = ?", 0).Order("id desc").
		Scan(&posts, "timeline = ?", timeline)
	r.Response.WriteJsonExit(posts)
}
