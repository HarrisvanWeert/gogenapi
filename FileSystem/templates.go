package filesystem

func getMainFileContent() string {
	return `package main

import (
	"fmt"
	"github.com/joho/godotenv"

)

func main() {
 err := godotenv.Load()
 if err != nil {
  fmt.Println("Error loading .env file")
  return
}	

}
`
}

func getHandlersFileContent() string {
	return `package handlers

import "github.com/gofiber/fiber/v2"

func HomeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello from Fiber!")
}
`
}

func GetDbPostgresFileContent() string {
	return `
	package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
}

	`
}

func GetDbSqlserverFileContent() string {
	return `
	package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}

	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}
}
`

}
