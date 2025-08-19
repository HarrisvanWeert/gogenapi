package directorysystem

import (
	"fmt"
	"os"
)

func MakeDir(name string) error {

	if name == "" {
		return fmt.Errorf("directory name cannot be empty")
	}

	err := os.MkdirAll(name, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	return nil
}
