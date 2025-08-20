package main

import (
	"fmt"
	"os"
	"os/exec"

	directorysystem "github.com/HarrisvanWeert/GoFiberCreate/DirectorySystem"
	filesystem "github.com/HarrisvanWeert/GoFiberCreate/FileSystem"
	"github.com/manifoldco/promptui"
)

func main() {

	gomod := ""
	fmt.Println("Enter the module name (e.g., github.com/username/project):")
	fmt.Scanln(&gomod)
	askUserforDb()

	paths := []string{"GOAPI", "GOAPI/handlers", "GOAPI/services", "GOAPI/db"}

	for _, path := range paths {
		err := directorysystem.MakeDir(path)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	err := filesystem.CreateAndWritetoTheFiles()
	if err != nil {
		fmt.Println("Error creating and writing to files:", err)
		return
	}

	err = runCommand("GOAPI", "go", "mod", "init", gomod)
	if err != nil {
		fmt.Println("Error initializing Go module:", err)
		return
	}

	// err = runCommand("GOAPI", "go", "get", "github.com/joho/godotenv")
	// if err != nil {
	// 	fmt.Println("Error getting godotenv package:", err)
	// 	return
	// }

	// err = runCommand("GOAPI", "go", "get", "github.com/gofiber/fiber/v2")
	// if err != nil {
	// 	fmt.Println("Error getting Fiber package:", err)
	// 	return
	// }
	err = runCommand("GOAPI", "go", "mod", "tidy")
	if err != nil {
		fmt.Println("Error tidying Go module:", err)
		return
	}

	fmt.Println("Go module initialized successfully.")
	fmt.Println("done with main function")

}

func runCommand(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func askUserforDb() {
	dbOptions := []string{"Postgres", "SqlServer"}

	prompt := promptui.Select{
		Label: "Select a database",
		Items: dbOptions,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	fmt.Println("You selected:", result)

	filesystem.MakeDbFile(result)
}
