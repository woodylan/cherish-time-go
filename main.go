package main

import (
	_ "cherish-time-go/routers"

	"github.com/astaxie/beego"
	"cherish-time-go/db"
	"cherish-time-go/cache"
)

//var MainRedis cache.Cache

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//连接数据库
	db, err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//连接redis
	cache.Init()

	beego.Run()
}
