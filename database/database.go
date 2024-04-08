package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	const MYSQL = "mr_cyp:L0rem!psum@tcp(127.0.0.1:3306)/go_fiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// const POSTGRES = "user=postgres password=postgres dbname=go_fiber_gorm port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	dsn := MYSQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	fmt.Println("Connected to database")
}
