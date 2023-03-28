package database

import (
	"sample_go/config"
	"sample_go/resources"
)

func GetAllUsers() []resources.UsersOutput {
	db := config.MariaDb

	var users []resources.UsersOutput

	if err := db.Debug().Table("users").Find(&users).Error; err != nil {
		return users
	}

	return users
}
