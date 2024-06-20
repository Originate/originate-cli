package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Originate/go-utilities/utils"
	"github.com/Originate/originate-cli/actions/migrations/templates"
)

type CreateMigrationInput struct {
	MigrationName string
	MigrationsDir string
}

var timestampFormat = "20060102150405"

func CreateNewMigration(i CreateMigrationInput) int {
	if i.MigrationName == "" {
		fmt.Println("Migration needs to have a name set using the --name flag")
		return 1
	}

	filename := fmt.Sprintf(
		"%v_%v.%v",
		time.Now().UTC().Format(timestampFormat),
		utils.ToSnakeCase(i.MigrationName),
		"sql",
	)

	path := filepath.Join(i.MigrationsDir, filename)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return 1
	}

	f, err := os.Create(path)
	if err != nil {
		return 1
	}
	defer f.Close()

	if err := templates.SQLMigrationTemplate.Execute(f, nil); err != nil {
		return 1
	}

	fmt.Printf("Created new file: %s\n", filename)
	return 0
}
