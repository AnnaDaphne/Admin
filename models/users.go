package models

type User struct {
    Id          int     `form:"-"`
    Username    string  `form:"username" valid:"Required;AlphaNumeric;MinSize(3);MaxSize(20)"`
    Password    string  `form:"password" valid:"Required;MinSize(3);MaxSize(100)"`
}

func (p *User) TableName() string {
    return "users"
}