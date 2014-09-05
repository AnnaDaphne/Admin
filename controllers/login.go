package controllers

import (

)

type LoginController struct {
	AdminController
}

func (this *LoginController) Get() {
	this.Data["Login"] = true
	this.Data["Title"] = "ACP Login"
	this.TplNames = "admin/pages/login.tpl"
}