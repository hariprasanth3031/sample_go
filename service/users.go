package service

import (
	"fmt"
	"sample_go/config"
	"sample_go/database"
	"sample_go/resources"
)

func GetAllUsers() resources.JsonOutput {

	fmt.Println("url is", config.Env.BaseUrl)

	output := database.GetAllUsers()
	return resources.JsonOutput{
		Code: 200,
		Data: output,
	}
}
