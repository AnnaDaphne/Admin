package controllers

import (
    "github.com/astaxie/beego"
)

type AdminController struct {
    beego.Controller
}

func (this *AdminController) Prepare() {
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
}