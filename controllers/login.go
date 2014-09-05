package controllers

import (
    "fmt"
    "code.google.com/p/go.crypto/bcrypt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
    models "github.com/annadaphne/Admin/models"
)

type LoginController struct {
    AdminController
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
            o := orm.NewOrm()
            user := models.User{Username: postUser.Username}
            err := o.Read(&user, "Username")

            if err == orm.ErrNoRows {
                flash.Error(userpwErr)
                beego.Debug(fmt.Sprintf("Incorrect username entered: ", postUser.Username))
            } else {
                err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(postUser.Password))
                if err == nil {
                    user.Lastip = IPToU32(this.Ctx.Input.IP())
                    if _, err = o.Update(&user); err != nil {
                        beego.Error(err)
                    }

                    beego.Debug(fmt.Sprintf("Login successful for user ID: ", user.Id))
                    this.Redirect(this.UrlFor("PlaceController.Get"), 302)
                    return
                }
                beego.Debug(fmt.Sprintf("Incorrect password entered: ", err))
                flash.Error(userpwErr)
            }
        }
    }

    flash.Store(&this.Controller)
    this.Redirect(this.UrlFor("LoginController.Get"), 302)
}

// Temp
func (this *LoginController) GenBcrypt() {
    beego.AutoRender = false
    password := []byte(this.Ctx.Input.Param(":plaintxt"))
    hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
    if err != nil {
        beego.Error(err)
    } else {
        this.Ctx.WriteString(string(hashedPassword))
    }
}