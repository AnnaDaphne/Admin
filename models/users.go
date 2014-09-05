package models

import (
    "time"
)

type User struct {
    Id          uint        `form:"-"`
    Username    string      `form:"username" valid:"Required;AlphaNumeric;MinSize(3);MaxSize(20)"`
    Password    string      `form:"password" valid:"Required;MinSize(3);MaxSize(100)"`
    Acl         uint8       `form:"-"`
    Lastlogin   time.Time   `orm:"auto_now;type(datetime)"`
    Lastip      uint32      `form:"-"`
}

func (p *User) TableName() string {
    return "users"
}