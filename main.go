package main

import (
    _ "github.com/annadaphne/Admin/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    /*
     * TODO: Config
     */
    orm.RegisterDriver("mysql", orm.DR_MySQL)
    orm.RegisterDataBase("default", "mysql", "root:123456@/annadaphne?charset=utf8")
}

func main() {
    beego.Run()
}

