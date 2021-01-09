package models

import (
	"github.com/spiderman930706/gin_admin/models"
)

type User struct {
	models.User
	Phone string `gorm:"type:varchar(15);unique_index" json:"phone" admin:"list:modified_time;type:time"`
}

func CheckAuth(username, password string) bool {
	var user User
	db.Select("id").Where(&User{User: models.User{Username: username, Password: password}}).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}