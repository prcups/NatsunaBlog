package info

import (
  "NatsunaBlog/app/dao"
  "NatsunaBlog/app/model"
  "github.com/gogf/gf/net/ghttp"
)

type InfoRes struct {
  Title string `json:title orm:title`
  desp string `json:desp orm:desp`
  Url string `json:url orm:url`
}

func GetInfo(r *ghttp.Request) {
  var info *model.DBBLOGCONFIG
  err := dao.DBBLOGCONFIG.Scan(info)
  if err != nil {
    r.Response.WritelnExit("GETINFO: " + err.Error())
  }
  r.Response.WriteJsonExit(info)
}

func GetNav(r *ghttp.Request) {
  nav, err := dao.DBBLOGONNAV.All()
  if err != nil {
    r.Response.WritelnExit("GETNAV: " + err.Error())
  }
  r.Response.WriteJsonExit(nav)
}