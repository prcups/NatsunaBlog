package check

import (
	"NatsunaBlog/app/dao"
	"github.com/gogf/gf/net/ghttp"
)

type checkRes struct {
	IsChecked bool 	`json:"isChecked"`
	User string 	`json:"user"`
}

func LogOut(r *ghttp.Request) {
	r.Session.Clear()
	r.Response.WriteExit("OK")
}

func Check(r *ghttp.Request) {
	if s := r.Session.GetString("user"); s != ""{
		r.Response.WriteJsonExit(checkRes{
			IsChecked: true,
			User: s,
		})
	} else if r.Get("username") == "" {
		r.Response.WriteJsonExit(checkRes{
			IsChecked: false,
			User: "",
		})
	} else if c, _ := dao.DBBLOGCONFIG.Count("username = ? AND password = ?", r.Get("username"), r.Get("password")); c > 0 {
		r.Session.Set("user", r.Get("username"))
		r.Response.WriteJsonExit(checkRes{
			IsChecked: true,
			User: r.Session.GetString("user"),
		})
	} else {
		r.Response.WriteJsonExit(checkRes{
			IsChecked: false,
			User: "",
		})
	}
}
