package main

import (
	_ "cherish-time-go/routers"

	"github.com/astaxie/beego"
	"cherish-time-go/db"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	beego.Run()
}
