package main

import (
	"fmt"

	"github.com/Marmar9/wolne-lektury-go/server/db/models"
	"github.com/Marmar9/wolne-lektury-go/server/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dbUser, dbPass, dbName := utils.GetEnv()
	connectionString := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Kind{}, &models.Author{}, &models.Lecture{})
}
