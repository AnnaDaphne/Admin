package routers

import (
    "github.com/mrsaints/AD-ACP/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/admin/place/", &controllers.PlaceController{})
    beego.Router("/admin/place/add", &controllers.PlaceController{}, "get,post:AddPlace")
}