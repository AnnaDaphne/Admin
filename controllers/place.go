package controllers

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    models "github.com/annadaphne/Admin/models"
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
        flash.Notice("No places found. This may be due to a server error.")
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
    flash := beego.NewFlash()
    serverErr := "We are unable to add the place at this time due to a server error. Please try again later."

    place := new(models.Place)
    if err := this.ParseForm(place); err != nil {
        flash.Error(serverErr)
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
            o := orm.NewOrm()
            id, err := o.Insert(place)

            if err == nil {
                flash.Notice("The place has been successfully added.")
                beego.Debug(fmt.Sprintf("Added a new place with ID: ", id))
            } else {
                flash.Error(serverErr)
                beego.Error(err)
            }

            flash.Store(&this.Controller)
            this.Redirect(this.UrlFor("PlaceController.AddPlace"), 302)
        }
    }
}