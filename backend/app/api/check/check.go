package check

import (
	"backend/app/dao"
	"github.com/gogf/gf/v2/net/ghttp"
)

type checkRes struct {
	IsChecked bool   `json:"isChecked"`
	User      string `json:"user"`
}

func LogOut(r *ghttp.Request) {
	_ = r.Session.RemoveAll()
	r.Response.WriteExit("OK")
}

func Check(r *ghttp.Request) {
	if s, _ := r.Session.Get("user"); s.String() != "" {
		r.Response.WriteJsonExit(checkRes{
			IsChecked: true,
			User:      s.String(),
		})
	} else if r.Get("username").String() == "" {
		r.Response.WriteJsonExit(checkRes{
			IsChecked: false,
			User:      "",
		})
	} else if c, _ := dao.DBBLOGCONFIG.Ctx(r.GetCtx()).Count("username = ? AND password = ?", r.Get("username"), r.Get("password")); c > 0 {
		_ = r.Session.Set("user", r.Get("username"))
		r.Response.WriteJsonExit(checkRes{
			IsChecked: true,
			User:      r.Get("username").String(),
		})
	} else {
		r.Response.WriteJsonExit(checkRes{
			IsChecked: false,
			User:      "",
		})
	}
}
