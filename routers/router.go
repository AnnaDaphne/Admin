package routers

import (
    "github.com/annadaphne/Admin/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/admin/logout/", &controllers.AdminController{}, "get:Logout")
    beego.Router("/admin/login/", &controllers.LoginController{})
    beego.Router("/admin/generate/:plaintxt", &controllers.AdminController{}, "get:GenBcrypt")
    beego.Router("/admin/place/", &controllers.PlaceController{})
    beego.Router("/admin/place/add", &controllers.PlaceController{}, "get,post:AddPlace")
}