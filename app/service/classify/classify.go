package classify

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

//获取所有分类
func GetAllClassify(r *ghttp.Request) {
	classify, err := dao.DBBLOGPOST.Where("hid = ?", 0).FindArray("distinct classify")
	if err != nil {
		r.Response.WritelnExit("GETALLCLASSIFY: " + err.Error())
	}
	r.Response.WriteJsonExit(classify)
}

//获取某个分类的文章
func GetPostsOfClassify(r *ghttp.Request) {
	classify := r.GetString("classify")
	var posts []GetPostsElement
	dao.DBBLOGPOST.Where("hid = ?", 0).Order("id desc").
		Scan(&posts, "classify = ?", classify)
	r.Response.WriteJsonExit(posts)
}
