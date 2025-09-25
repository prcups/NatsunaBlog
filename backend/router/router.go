package router

import (
	"blog/app/api/check"
	"blog/app/api/image"
	"blog/app/service/classify"
	"blog/app/service/middleware"
	"blog/app/service/post"
	"blog/app/service/timeline"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/ava-api", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		group.GET("/get-posts", post.GetPosts)
		group.GET("/get-page-num", post.GetPageNum)
		group.Group("/post", func(group *ghttp.RouterGroup) {
			group.POST("/update", post.UpdatePost)
			group.POST("/delete", post.DeletePost)
			group.GET("/get-one", post.GetOnePost)
		})
		group.Group("/image", func(group *ghttp.RouterGroup) {
			group.POST("/upload", image.UploadImage)
		})
		group.Group("/get-classify", func(group *ghttp.RouterGroup) {
			group.GET("/get-all", classify.GetAllClassify)
			group.GET("/get-posts", classify.GetPostsOfClassify)
		})
		group.Group("/get-timeline", func(group *ghttp.RouterGroup) {
			group.GET("/get-all", timeline.GetAllTimeLine)
			group.GET("/get-posts", timeline.GetPostsOfTimeLine)
		})
		group.POST("/check", check.Check)
		group.GET("/logout", check.LogOut)
	})
}
