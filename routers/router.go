package routers

import (
    "github.com/mrsaints/AD-ACP/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/admin", &controllers.AdminController{})
    beego.Router("/admin/place/create", &controllers.AdminController{}, "get:GetCreatePlace")
    beego.Router("/admin/place/create", &controllers.AdminController{}, "post:PostCreatePlace")
}
