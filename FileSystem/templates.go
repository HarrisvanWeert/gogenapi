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
