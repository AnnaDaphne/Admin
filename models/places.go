package models

import (
    "github.com/astaxie/beego/orm"
    //_ "github.com/go-sql-driver/mysql"
)

type Place struct {
    Id          int     `form:"-"`
    Name        string  `form:"name"`
    Description string  `form:"description"`
    Website     string  `form:"website"`
    Lat         string  `form:"gpsLat" orm:type(POINT)`
    Long        string  `form:"gpsLong" orm:type(POINT)`
    Type        int     `form:"-"`
}

func (p *Place) TableName() string {
    return "places"
}

func init() {
    orm.RegisterModel(new(Place))
}
