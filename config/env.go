package config

import "os"

var Env *EnvData

type EnvData struct {
	Dbcon string
}

func InitializebEnv() {
	Env = &EnvData{
		Dbcon: os.Getenv("Dbcon"),
	}

}
