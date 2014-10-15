// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beerestangular/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/auth_user",
			beego.NSInclude(
				&controllers.AuthUserController{},
			),
		),

		beego.NSNamespace("/barang",
			beego.NSInclude(
				&controllers.BarangController{},
			),
		),

	)
	beego.AddNamespace(ns)
	beego.SetStaticPath("/bower_components","client/bower_components")
	beego.SetStaticPath("/css","client/css")
	beego.SetStaticPath("/js","client/js")
	beego.SetStaticPath("/views","client/views")
	beego.Router("/", &controllers.MainController{})
}
