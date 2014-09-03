package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    models "github.com/mrsaints/AD-ACP/models"
)

type AdminController struct {
    beego.Controller
}

func (this *AdminController) Get() {
    this.Layout = "admin/layout.tpl"
    this.TplNames = "admin/index.tpl"

    this.Data["Title"] = "Manage Places"

    o := orm.NewOrm()
    var places []*models.Place
    num, err := o.Raw("SELECT `id`, `name`, `description`, `website`, X(GeomFromText(AsText(`gps`))) as `long`, Y(GeomFromText(AsText(`gps`))) as `lat` FROM `places`").QueryRows(&places)
    //num, err := o.QueryTable("places").All(&places)
    if err != orm.ErrNoRows && num > 0 {
        this.Data["places"] = places
    } else {
        beego.Debug(err)
    }
}

func (this *AdminController) GetCreatePlace() {
    this.Layout = "admin/layout.tpl"
    this.TplNames = "admin/create_place.tpl"

    this.Data["Title"] = "Create Place"
}

func (this *AdminController) PostCreatePlace() {
    o := orm.NewOrm()
    place := models.Place{}
    if err := this.ParseForm(&place); err != nil {
        beego.Error("Could not parse the form. Reason: ", err)
    } else {
        var r orm.RawSeter
        r = o.Raw("INSERT INTO `places` (`name`, `description`, `website`, `gps`) VALUES (?,?,?,PointFromText(?))")
        id, err :=  r.SetArgs(place.Name, place.Description, place.Website, "POINT(" + place.Long + " " + place.Lat + ")").Exec()
        if err == nil {
            beego.Debug(id)
        } else {
            beego.Debug(err)
        }
        this.Redirect("/admin/place/create",302)
    }
}