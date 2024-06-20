package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Originate/originate-cli/actions/api"
	"github.com/Originate/originate-cli/actions/migrations"
	"github.com/Originate/originate-cli/utils"
)

func main() {
	help := flag.Bool("help", false, "Prints out the help instructions")
	name := flag.String("name", "", "Name of the resource to be generated")
	org := flag.String("org", "Originate", "The organization to be used in the go.mod module name")
	migrationsDir := flag.String("migrations-dir", "database/migrations", "The relative path to the database migrations folder, defaults to \"database/migrations\"")

	// This needs to be below all the flag definitions
	utils.ParseFlags()

	if *help {
		fmt.Println("The CLI currently supports generating APIs and managing migrations, check the Github README for more info, here is the usage for the flags:")
		flag.Usage()
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	entity := flag.Arg(0)
	action := flag.Arg(1)

	actionAndEntity := strings.Join([]string{entity, action}, " ")
	switch actionAndEntity {
	case "api new":
		os.Exit(
			api.GenerateNewAPI(api.GenerateAPIConfig{
				Name:         *name,
				Organization: *org,
			}),
		)
	case "migrations new":
		os.Exit(
			migrations.CreateNewMigration(migrations.CreateMigrationInput{
				MigrationName: *name,
				MigrationsDir: *migrationsDir,
			}),
		)
	case "migrations up":
		os.Exit(
			migrations.Up(migrations.MigrationsInput{
				MigrationsDir: *migrationsDir,
				Context:       context.Background(),
			}),
		)
	case "migrations down":
		os.Exit(
			migrations.Down(migrations.MigrationsInput{
				MigrationsDir: *migrationsDir,
				Context:       context.Background(),
			}),
		)
	case "migrations reset":
		os.Exit(
			migrations.Reset(migrations.MigrationsInput{
				MigrationsDir: *migrationsDir,
				Context:       context.Background(),
			}),
		)
	default:
		fmt.Printf("unrecognized action: %s", actionAndEntity)
		os.Exit(1)
	}
}
