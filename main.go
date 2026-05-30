
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Database connection
	dsn := "host=localhost user=postgres password=password dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	

}