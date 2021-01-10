package models

import (
	"github.com/spiderman930706/gin_admin/models"
)

type User struct {
	models.User
	Phone string `gorm:"column:phone_num;type:varchar(15);unique_index" json:"phone"`
}

func CheckAuth(username, password string) bool {
	var user User
	db.Select("id").Where(&User{User: models.User{Username: username, Password: password}}).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}
