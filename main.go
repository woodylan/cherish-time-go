package main

import (
	_ "cherish-time-go/routers"

	"github.com/astaxie/beego"
	"cherish-time-go/models/init"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	initModel.Init()
	beego.Run()
}
