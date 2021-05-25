package classify

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

//获取所有分类
func GetAllClassify(r *ghttp.Request) {
	classify, err := dao.DBBLOGPOST.FindArray("classify")
	if err != nil {
		r.Response.WritelnExit("GETALLCLASSIFY: " + err.Error())
	}
	r.Response.WriteJsonExit(classify)
}

//获取某个分类的文章（分页）
func GetPostOfClassify(r *ghttp.Request) {
	classify := r.GetString("classify")
	page := r.GetInt("page")
	var posts GetPostsElement
	dao.DBBLOGPOST.Order("ontop desc, id desc").Limit((page-1)*post.PostsInOnePage, post.PostsInOnePage).
		Scan(&posts, "classify = ?", classify)
	r.Response.WriteJsonExit(posts)
}
