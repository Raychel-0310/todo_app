package main

import (
	"fmt"
	"go_rest_api/db"
	"go_rest_api/model"
)

// % GO_ENV=dev go run migrate/migrate.go
func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}