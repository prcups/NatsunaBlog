package timeline

import (
	"NatsunaBlog/app/dao"
	"github.com/gogf/gf/net/ghttp"
)

type GetPostsElement struct {
	Id int			`json:"id"`
	Time string	`json:"time"`
	Title string	`json:"title"`
	OnTop bool `json:"ontop"`
	Tag string `json:"tag"`
	Classify string `json:"classify"`
}

//获取所有归档
func GetAllTimeLine(r *ghttp.Request) {
	timeline, err := dao.DBBLOGPOST.FindArray("distinct timeline")
	if err != nil {
		r.Response.WritelnExit("GETALLTIMELINE: " + err.Error())
	}
	r.Response.WriteJsonExit(timeline)
}

//获取单个归档下的内容
func GetPostsOfTimeLine(r *ghttp.Request) {
	timeline := r.GetString("timeline")
	var posts []GetPostsElement
	dao.DBBLOGPOST.Order("id desc").
		Scan(&posts, "timeline = ?", timeline)
	r.Response.WriteJsonExit(posts)
}
