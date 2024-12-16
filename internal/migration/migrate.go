package main

import (
	"fmt"

	"github.com/aminasadiam/Shop/internal/database"
	"github.com/aminasadiam/Shop/models"
)

func main() {
	db, err := database.InitDB("root:19045522@tcp(localhost:3306)/shop_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})

	fmt.Println("Database migration completed successfully")
}
