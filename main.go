package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Task struct {
	gorm.Model
	Title    string
	Priority string
}

// Session Configuration
type Session struct {
	DryRun                   bool
	PrepareStmt              bool
	NewDB                    bool
	SkipHooks                bool
	SkipDefaultTransaction   bool
	DisableNestedTransaction bool
	AllowGlobalUpdate        bool
	FullSaveAssociations     bool
	QueryFields              bool
	CreateBatchSize          int
	Context                  context.Context
	Logger                   logger.Interface
	NowFunc                  func() time.Time
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

	CheckStatement()
}

func CheckStatement() {

	var task Task

	fmt.Printf("%+v\n", task)
}

func RetriveRecords() {
	tasks := []Task{}
	results := db.Find(&tasks)
	// SELECT * FROM Task

	fmt.Println(results)
	fmt.Println(tasks)

}

func UpdateRecord() {
	var task Task

	db.First(&task)

	task.Title = "Update Title"

	db.Save(task)
	// UPDATE All column

	db.First(&task, task.ID)
	// SELECT * FROM task WHERE ID = XX

	fmt.Printf("%+v\n", task)
}

func CreateRecord() {
	task := Task{
		Title:    "Create",
		Priority: "A",
	}

	db.Create(&task)

	fmt.Printf("%+v\n", task)
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

	fmt.Printf("%+v\n", tasks)
}
