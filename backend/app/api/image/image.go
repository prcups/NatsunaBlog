package image

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
)

type UploadImageRes struct {
	Name string `json:"name"`
}

var imageSaveUrl = "static/upload_images"

func UploadImage(r *ghttp.Request) {
	if t, _ := r.Session.Get("user"); t.String() == "" {
		r.Response.WriteExit("Not Login")
	}
	file := r.GetUploadFile("image")
	if file == nil {
		r.Response.WritelnExit("No File to Upload")
	}
	saveFileName, err := file.Save(imageSaveUrl, true)
	if err != nil {
		log.Fatal(err)
	}
	r.Response.WriteJsonExit(UploadImageRes{
		Name: saveFileName,
	})
}
