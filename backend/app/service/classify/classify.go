package classify

import (
	"backend/app/dao"
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

// 获取所有分类
func GetAllClassify(r *ghttp.Request) {
	classify, err := dao.DBBLOGPOST.Ctx(r.GetCtx()).Where("hid = ?", 0).Array("distinct classify")
	if err != nil {
		r.Response.WritelnExit("GETALLCLASSIFY: " + err.Error())
	}
	r.Response.WriteJsonExit(classify)
}

// 获取某个分类的文章
func GetPostsOfClassify(r *ghttp.Request) {
	classify := r.Get("classify").String()
	var posts []GetPostsElement
	dao.DBBLOGPOST.Ctx(r.GetCtx()).Where("hid = ?", 0).Order("id desc").
		Scan(&posts, "classify = ?", classify)
	r.Response.WriteJsonExit(posts)
}
