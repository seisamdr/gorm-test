package main

import (
	"fmt"

	//package used to read the .env file
	_ "github.com/lib/pq"

	"gorm/models"
	model "gorm/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func main() {
	var err error
	// DB, err = gorm.Open("mysql", "root:@/egommerce?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open("postgres", "postgresql://postgres:root@127.0.0.1/postgres?sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer DB.Close()

	fmt.Println("Connection Opened to Database")

	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&model.Student{})

	data := models.Student{}
	if DB.Find(&data).RecordNotFound() {
		println("================ Seeding Data ================")
		seederUser()
	}
}

func seederUser() {
	data := models.Student{
		Student_id:      1,
		Student_name:    "John",
		Student_age:     20,
		Student_address: "New York",
		Student_phone:   "1234567890",
	}
	DB.Create(&data)
}