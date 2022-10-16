package middleware

import "github.com/gogf/gf/v2/net/ghttp"

func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
