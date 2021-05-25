package timeline

import (
	"NatsunaBlog/app/dao"
	"NatsunaBlog/app/service/post"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

type GetPostsElement struct {
	Id int			`json:"id"`
	Time gtime.Time	`json:"time"`
	Title string	`json:"title"`
	OnTop bool `json:"ontop"`
	Tag string `json:"tag"`
	Classify string `json:"classify"`
}

//获取所有归档
func GetAllTimeLine(r *ghttp.Request) {
	timeline, err := dao.DBBLOGPOST.FindArray("timeline")
	if err != nil {
		r.Response.WritelnExit("GETALLTIMELINE: " + err.Error())
	}
	r.Response.WriteJsonExit(timeline)
}

//获取单个归档下的内容
func GetPostOfTimeLine(r *ghttp.Request) {
	timeline := r.GetString("timeline")
	page := r.GetInt("page")
	var posts GetPostsElement
	dao.DBBLOGPOST.Order("ontop desc, id desc").Limit((page-1)*post.PostsInOnePage, post.PostsInOnePage).
		Scan(&posts, "timeline = ?", timeline)
	r.Response.WriteJsonExit(posts)
}
