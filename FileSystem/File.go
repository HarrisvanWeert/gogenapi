package filesystem

import (
	"fmt"
	"os"
)

func CreateAndWritetoTheFiles() error {
	// Create main.go
	mainfile, err := Makefile("main", "GOAPI")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer mainfile.Close()

	_, err = mainfile.WriteString(`package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
`)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	// Create handlers.go
	handlersfile, err := Makefile("handlers", "GOAPI/handlers")
	if err != nil {
		fmt.Println("Error creating handlers file:", err)
		return err
	}
	defer handlersfile.Close()

	_, err = handlersfile.WriteString(`package handlers

import "github.com/gofiber/fiber/v2"

func HomeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello from Fiber!")
}
`)
	if err != nil {
		fmt.Println("Error writing to handlers file:", err)
		return err
	}

	// Create db.go
	dbfile, err := Makefile("db", "GOAPI/db")
	if err != nil {
		fmt.Println("Error creating db file:", err)
		return err
	}
	defer dbfile.Close()

	_, err = dbfile.WriteString(`package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the PostgreSQL database connection
func InitDB(host, user, password, dbname, port, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db
	log.Println("PostgreSQL connection established successfully")
	return DB, nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
`)
	if err != nil {
		fmt.Println("Error writing to db file:", err)
		return err
	}

	return nil
}

func Makefile(name string, location string) (*os.File, error) {
	if name == "" || location == "" {
		return nil, fmt.Errorf("name and location cannot be empty")
	}

	// Ensure directory exists
	if err := os.MkdirAll(location, 0755); err != nil {
		return nil, err
	}

	file, err := os.Create(location + "/" + name + ".go")
	if err != nil {
		return nil, err
	}

	return file, nil
}
