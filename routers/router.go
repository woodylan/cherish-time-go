// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"cherish-time-go/controllers"

	"github.com/astaxie/beego"
	"cherish-time-go/controllers/weapp/time"
)

func init() {
	ns := beego.NewNamespace("/api/weapp/v1/",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/time",
			beego.NSRouter(
				"/detail",&timeComtroller.TimeDetailController{},"*:Detail",
			),
			beego.NSRouter(
				"/list",&timeComtroller.TimeListController{},"*:List",
			),
		),
	)
	beego.AddNamespace(ns)
}
