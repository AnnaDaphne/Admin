package controllers

import (
    //"fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    models "github.com/mrsaints/AD-ACP/models"
)

type PlaceController struct {
    AdminController
}

func (this *PlaceController) Get() {
    this.Data["Title"] = "View Places"
    this.TplNames = "admin/pages/view_places.tpl"

    o := orm.NewOrm()
    var places []*models.Place
    num, err := o.QueryTable("Places").All(&places)

    if err != orm.ErrNoRows && num > 0 {
        this.Data["Places"] = places
    } else {
        flash := beego.NewFlash()
        flash.Notice("No places found. There may be an error on the server.")
        flash.Store(&this.Controller)
        beego.Error(err)
    }
}

func (this *PlaceController) AddPlace() {
    this.Data["Title"] = "Add Place"
    this.TplNames = "admin/pages/add_place.tpl"

    beego.ReadFromRequest(&this.Controller)

    if this.Ctx.Input.IsPost() {
        this._AddPlace()
    }
}

func (this *PlaceController) _AddPlace() {
    place := new(models.Place)

    if err := this.ParseForm(place); err != nil {
        beego.Error("Could not parse form values to struct: ", err)
    } else {
        valid := validation.Validation{}
        isValid, _ := valid.Valid(place)

        if !isValid {
            // Store form session
            this.Data["Place"] = place
            this.Data["Errors"] = valid.ErrorsMap
            beego.Debug(this.Data["Errors"])
        } else {
            flash := beego.NewFlash()
            o := orm.NewOrm()
            _, err := o.Insert(place)

            if err == nil {
                flash.Notice("The place has been successfully added.")
            } else {
                flash.Error("We are unable to add the place at this time due to a database error. Please try again later.")
                beego.Error(err)
            }

            flash.Store(&this.Controller)
            this.Redirect(this.UrlFor("PlaceController.AddPlace"), 302)
        }
    }
}