package config

import (
	"fmt"

	"../structs"
	"github.com/jinzhu/gorm"
)

const (
	DATABASE_USERNAME = "root"
	DATABASE_PASSWORD = "rootpw"
	DATABASE_HOST     = "127.0.0.1"
	DATABASE_PORT     = "3306"
	DATABASE_NAME     = "parking_slots"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	templateDsn := "%s:%s@tcp(%s:%s)/"
	dsn := fmt.Sprintf(
		templateDsn,
		DATABASE_USERNAME,
		DATABASE_PASSWORD,
		DATABASE_HOST,
		DATABASE_PORT,
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + DATABASE_NAME + ";")

	db, err = gorm.Open("mysql", dsn+DATABASE_NAME+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(structs.Slot{})
	return db
}
