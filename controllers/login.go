package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/validation"
    "github.com/annadaphne/Admin/modules/auth"
    models "github.com/annadaphne/Admin/models"
)

type LoginController struct {
    AdminController
}

func (this *LoginController) Prepare() {
    this.AdminController.Prepare()

    if this.LoggedIn {
        this.Redirect(this.UrlFor("PlaceController.Get"), 302)
        return
    }
}

func (this *LoginController) Get() {
    this.Data["Login"] = true
    this.Data["Title"] = "ACP Login"
    this.TplNames = "admin/pages/login.tpl"
    beego.ReadFromRequest(&this.Controller)
}

func (this *LoginController) Post() {
    flash := beego.NewFlash()
    serverErr := "We are unable to process your request at this time due to a server error. Please try again later."
    userpwErr := "The email or password you entered is incorrect."

    postUser := new(models.User)
    if err := this.ParseForm(postUser); err != nil {
        flash.Error(serverErr)
        beego.Error("Could not parse form values to struct: ", err)
    } else {
        valid := validation.Validation{}
        isValid, _ := valid.Valid(postUser)

        if !isValid {
            flash.Error(userpwErr)
            this.Data["Errors"] = valid.ErrorsMap
            beego.Debug(this.Data["Errors"])
        } else {
            user := new(models.User)
            if auth.VerifyUser(user, postUser.Username, postUser.Password) {
                this.CompleteLogin(user)
                this.Redirect(this.UrlFor("PlaceController.Get"), 302)
                return
            } else {
                flash.Error(userpwErr)
            }
        }
    }

    flash.Store(&this.Controller)
    this.Redirect(this.UrlFor("LoginController.Get"), 302)
}