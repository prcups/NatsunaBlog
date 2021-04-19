package main

import (
	_ "NatsunaBlog/boot"
	_ "NatsunaBlog/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
