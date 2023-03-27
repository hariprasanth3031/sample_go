package config

import (
	mysql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	"gorm.io/gorm"
)

var MariaDb *gorm.DB

func InitializeDb() {

	db, err := gorm.Open(mysql.Open(Env.Dbcon+"?&parseTime=True"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	MariaDb = db

}
