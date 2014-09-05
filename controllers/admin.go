package controllers

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "github.com/annadaphne/Admin/modules/auth"
    "github.com/annadaphne/Admin/modules/utils"
    models "github.com/annadaphne/Admin/models"
)

type AdminController struct {
    beego.Controller
    User              *models.User
    LoggedIn          bool
}

func (this *AdminController) Prepare() {
    beego.AutoRender = true
    this.Layout = "admin/layout.tpl"

    this.Data["HeadStyles"] = []string{
        "//cdn.jsdelivr.net/pure/0.5.0/pure-min.css",
        "//yui.yahooapis.com/pure/0.5.0/grids-responsive-min.css",
        "//fonts.googleapis.com/css?family=Roboto:400,500,300",
        "/static/css/admin.css",
    }

    this.Data["HeadScripts"] = []string{
        "//cdn.jsdelivr.net/modernizr/2.8.3/modernizr.min.js",
    }

    this.Data["Token"] = this.XsrfToken()
    utils.SecureHeaders(this.Ctx)

    if this.LoadUserSession() && this.User.Acl == 1 {
        this.LoggedIn = true
        beego.Debug(fmt.Sprintf("Loaded user from session with ID: ", this.User.Id))
    } else {
        loginURL := this.UrlFor("LoginController.Get")

        if loginURL != this.Ctx.Input.Uri() {
            this.Redirect(loginURL, 302)
            return
        }
    }
}

func (this *AdminController) CompleteLogin(user *models.User) {
    user.Lastip = utils.IPToU32(this.Ctx.Input.IP())
    o := orm.NewOrm()
    if _, err := o.Update(user); err != nil {
        beego.Error(err)
    }
    this.SessionRegenerateID()
    this.SetSession("uid", user.Id)
    beego.Debug(fmt.Sprintf("Login successful for user ID: ", user.Id))
}

func (this *AdminController) CompleteLogout() {
    beego.Debug("Logout successful.")
    this.DelSession("uid")
    this.DestroySession()
}

func (this *AdminController) Logout() {
    this.CompleteLogout()
    this.Redirect(this.UrlFor("LoginController.Get"), 302)
}

func (this *AdminController) LoadUserSession() bool {
    id, ok := this.GetSession("uid").(uint);
    if ok && id > 0 {
        o := orm.NewOrm()
        user := new(models.User)
        user.Id = id
        err := o.Read(user)
        if err == nil {
            this.User = user
            return true
        }
        this.CompleteLogout()
        beego.Error(err)
    }
    return false
}

// Temp
func (this *AdminController) GenBcrypt() {
    beego.AutoRender = false

    password := auth.HashPassword(this.Ctx.Input.Param(":plaintxt"))
    this.Ctx.WriteString(password)
}