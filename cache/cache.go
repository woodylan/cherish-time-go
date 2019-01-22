package cache

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
	)

var Bm cache.Cache

func Init() {
	var err error

	key := beego.AppConfig.String("redis_key")
	host := beego.AppConfig.String("redis_host")
	port := beego.AppConfig.String("redis_port")
	password := beego.AppConfig.String("redis_password")

	//config:=`{"key":"cherishTime","conn":"127.0.0.1:6379","dbNum":"0","password":""}`
	config := `{"key":"` + key + `","conn":"` + host + `:` + port + `","dbNum":"0","password":"` + password + `"}`

	Bm, err = cache.NewCache("redis", config)
	if err != nil {
		beego.Debug("Redis init fail")
	} else {
		beego.Debug("Redis init success")
	}
}
