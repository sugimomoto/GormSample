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

var db *gorm.DB

func main() {
	var err error

	dsn := "root:Password!@tcp(127.0.0.1:3306)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	RetriveRecords()
}

func RetriveRecords() {
	tasks := []Task{}
	results := db.Find(&tasks)
	// SELECT * FROM Task

	fmt.Println(results)
	fmt.Println(tasks)

}

func UpdateRecord() {

}

func CreateRecord() {
	task := Task{
		Title:    "Create",
		Priority: "A",
	}

	db.Create(&task)

	fmt.Println(task)
}

func CreateRecords() {

	tasks := []Task{
		{
			Title:    "Batch Insert 1",
			Priority: "A",
		},
		{
			Title:    "Batch Insert 2",
			Priority: "B",
		},
		{
			Title:    "Batch Insert 3",
			Priority: "C",
		},
	}

	db.Create(&tasks)

	fmt.Println(tasks)
}
