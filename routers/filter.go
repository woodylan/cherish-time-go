package routers

import (
	"github.com/astaxie/beego/context"
	_ "github.com/astaxie/beego/cache/redis"
	"cherish-time-go/cache"
	"cherish-time-go/logic/weapp/account"
	"cherish-time-go/modules/util"
	"cherish-time-go/global"
	"cherish-time-go/define/retcode"
)

var filterLoggedInUser = func(ctx *context.Context) {
	auth := ctx.Input.Header("Auth")
	if auth != "" {
		authData := accountLogic.AuthData{}

		redis := cache.Bm.Get(auth)
		if redis == nil {
			util.ThrowApi(ctx, retcode.ERR_NO_LOGIN, "用户未登录")
			return
		}

		util.JsonDecode(string(redis.([]byte)), &authData)

		//存到全局
		global.LoginUserInfo = &authData
	} else {
		util.ThrowApi(ctx, retcode.ERR_NO_LOGIN, "用户未登录")
		return
	}
}
