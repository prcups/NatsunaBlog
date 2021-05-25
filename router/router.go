package router

import (
	"NatsunaBlog/app/api/check"
	"NatsunaBlog/app/api/image"
	"NatsunaBlog/app/service/classify"
	"NatsunaBlog/app/service/middleware"
	"NatsunaBlog/app/service/post"
	"NatsunaBlog/app/service/timeline"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.AddStaticPath("/", "/public")
	s.AddStaticPath("/static", "/static")
	s.Group("/ava-api",func(group *ghttp.RouterGroup){
		group.Middleware(middleware.CORS)
		group.GET("/get-posts", post.GetPosts)
		group.GET("/get-page-num", post.GetPageNum)
		group.Group("/post", func(group *ghttp.RouterGroup){
			group.POST("/update", post.UpdatePost)
			group.POST("/delete", post.DeletePost)
			group.GET("/get-one", post.GetOnePost)
		})
		group.Group("/image", func(group *ghttp.RouterGroup){
			group.POST("/upload", image.UploadImage)
			group.POST("/delete", image.DeleteImage)
		})
		group.Group("/get-classify", func(group *ghttp.RouterGroup){
			group.GET("/get-all", classify.GetAllClassify)
			group.GET("/get-post", classify.GetPostOfClassify)
		})
		group.Group("/get-timeline", func(group *ghttp.RouterGroup){
			group.GET("/get-all", timeline.GetAllTimeLine)
			group.GET("/get-post", timeline.GetPostOfTimeLine)
		})
		group.POST("/check", check.Check)
		group.GET("/logout", check.LogOut)
	})
}
