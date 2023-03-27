package service

import (
	"sample_go/database"
	"sample_go/resources"
)

func GetAllUsers() resources.JsonOutput {

	output := database.GetAllUsers()
	return resources.JsonOutput{
		Code: 200,
		Data: output,
	}
}
