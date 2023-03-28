package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var Env *EnvData

type EnvData struct {
	Dbcon   string `required:"true"`
	BaseUrl string `required:"true"`
}

func InitializeEnv() {
	_ = godotenv.Load("Config.env")
	Env = &EnvData{
		Dbcon:   os.Getenv("Dbcon"),
		BaseUrl: os.Getenv("Baseurl"),
	}

	v := validator.New()
	err := v.Struct(*Env)
	if err != nil {
		fmt.Println(err)
		return
	}
	return

}
