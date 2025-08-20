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

	_, err = mainfile.WriteString(getMainFileContent())
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

	_, err = handlersfile.WriteString(getHandlersFileContent())
	if err != nil {
		fmt.Println("Error writing to handlers file:", err)
		return err
	}

	dotenvfile, err := MakeDotEnvFile("GOAPI")
	if err != nil {
		fmt.Println("Error creating .env file:", err)
	}

	_, err = dotenvfile.WriteString(`DB_HOST=` + "\n")
	if err != nil {
		fmt.Println("Error writing to .env file:", err)
	}

	return nil
}

func Makefile(name string, location string) (*os.File, error) {
	if name == "" || location == "" {
		return nil, fmt.Errorf("name and location cannot be empty")
	}

	if err := os.MkdirAll(location, 0755); err != nil {
		return nil, err
	}

	file, err := os.Create(location + "/" + name + ".go")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func MakeDotEnvFile(location string) (*os.File, error) {
	if location == "" {
		return nil, fmt.Errorf("location cannot be empty")
	}

	if err := os.MkdirAll(location, 0755); err != nil {
		return nil, err
	}

	file, err := os.Create(location + "/.env")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func MakeDbFile(dbchoice string) error {
	switch dbchoice {
	case "Postgres":
		fmt.Println("Creating Postgres DB file")
		dbfile, err := Makefile("db", "GOAPI/db")
		if err != nil {
			fmt.Println("Error creating db file:", err)
			return err
		}

		defer dbfile.Close()

		_, err = dbfile.WriteString(GetDbPostgresFileContent())
		if err != nil {
			fmt.Println("Error writing to db file:", err)
			return err
		}

	case "SqlServer":
		fmt.Println("Creating SqlServer DB file")
		dbfile, err := Makefile("db", "GOAPI/db")
		if err != nil {
			fmt.Println("Error creating db file:", err)
			return err
		}

		defer dbfile.Close()

		_, err = dbfile.WriteString(GetDbSqlserverFileContent())
		if err != nil {
			fmt.Println("Error writing to db file:", err)
			return err
		}

	default:
		fmt.Println("Unknown database choice")
	}

	return nil
}
