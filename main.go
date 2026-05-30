package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// Database connection
	dsn := "host=localhost user=postgres password=0909 dbname=Rag-System port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("err happen : database is not connect ", err)
	}
	DB = db

	app := fiber.New()
	// app.Use(recover.New())
	// app.Use(logger.New())

	fmt.Print("Hellow world My program is running")

}
