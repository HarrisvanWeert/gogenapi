package main

import (
	"fmt"
	"os"
	"os/exec"

	directorysystem "github.com/HarrisvanWeert/GoFiberCreate/DirectorySystem"
	filesystem "github.com/HarrisvanWeert/GoFiberCreate/FileSystem"
)

func main() {

	gomod := ""
	fmt.Println("Enter the module name (e.g., github.com/username/project):")
	fmt.Scanln(&gomod)

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

	err = runCommand("GOAPI", "go", "get", "github.com/gofiber/fiber/v2")
	if err != nil {
		fmt.Println("Error getting Fiber package:", err)
		return
	}
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
