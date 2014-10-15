package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["beerestangular/controllers:BarangController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:BarangController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beerestangular/controllers:BarangController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:BarangController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beerestangular/controllers:BarangController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:BarangController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beerestangular/controllers:BarangController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:BarangController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beerestangular/controllers:BarangController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:BarangController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["beerestangular/controllers:AuthUserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
