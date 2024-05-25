package main

import (
	"fmt"
	"github.com/ko44d/go-rest-api/db"
	"github.com/ko44d/go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
