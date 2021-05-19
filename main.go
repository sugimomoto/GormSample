package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title    string
	Priority string
}

func main() {
	dsn := "root:Password!@tcp(127.0.0.1:3306)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&Task{})

	db.Create(&Task{Title: "Hello", Priority: "A"})

	var task Task

	db.First(&task, 1)

	fmt.Println(task)

}
