package image

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"log"
)

type UploadImageRes struct{
	Name string `json:"name"`
}

var imageSaveUrl = "static/upload_images"

func UploadImage(r *ghttp.Request) {
	if r.Session.GetString("user") == "" {
		r.Response.WriteExit("Not Login")
	}
	file := r.GetUploadFile("image")
	if file == nil {
		r.Response.WritelnExit("No File to Upload")
	}
	saveFileName, err := file.Save(imageSaveUrl, true)
	if err != nil{
		log.Fatal(err)
	}
	r.Response.WriteJsonExit(UploadImageRes{
		Name: saveFileName,
	})
}

func DeleteImage(r *ghttp.Request) {
	if r.Session.GetString("user") == "" {
		r.Response.WriteExit("Not Login")
	}
	filePath := imageSaveUrl + "/" + gfile.Basename(r.GetString("url"))
	if gfile.Exists(filePath) {
		err := gfile.Remove(filePath)
		if err != nil {
			log.Fatal(err)
		}
	}
	r.Response.WritelnExit("OK")
}