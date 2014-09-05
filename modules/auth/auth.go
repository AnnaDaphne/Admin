package auth

import (
	"fmt"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "code.google.com/p/go.crypto/bcrypt"
    models "github.com/annadaphne/Admin/models"
)

func VerifyUsername(user *models.User, username string) bool {
	o := orm.NewOrm()
	user.Username = username
	err := o.Read(user, "Username")
	if err == nil {
		return true
	}
	beego.Debug(fmt.Sprintf("Incorrect username entered: ", username))
	return false
}

func VerifyPassword(plaintext string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plaintext))
	if err == nil {
		return true
	}
	beego.Debug(fmt.Sprintf("Incorrect password entered: ", err))
	return false
}

func VerifyUser(user *models.User, username string, password string) bool {
	if VerifyUsername(user, username) &&
		VerifyPassword(password, user.Password) {
		return true
	}
	return false
}

func HashPassword(plaintext string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintext), 10)
	if err != nil {
		return string(hashedPassword)
	}
	beego.Error(err)
	return ""
}