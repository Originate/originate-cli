package templates

import (
	"text/template"

	"github.com/nao1215/markdown"
)

// Had to build the README using markdown because we can't escape backticks inside backticks
var READMEContent = markdown.NewMarkdown(nil).
	H1("{{.FormattedName}}\n").
	H2("Dependencies\n").
	PlainText("Install the following tools for all the utilities to work:\n").
	BulletList(
		"[goose](https://github.com/pressly/goose?tab=readme-ov-file#install) for migrations",
		"[golangci-lint](https://golangci-lint.run/welcome/install/#local-installation) for linting",
		"[make](https://www.gnu.org/software/make/) for easier run of common commands\n\t\t- Most of the Unix-based OS's already have make installed, check by running `make --version`\n",
	).
	H2("Running Instructions\n").
	PlainText("The `config/config.yml` is a config template for each app, check if the values match your local config before running the app. It is set up to run the app in the 3000 port and to connect to a PostgreSQL database set up with the default [docker](https://hub.docker.com/_/postgres) instructions, check the 'How to use this image' section for reference.\n").
	PlainText("The `--config-file` flag is used to specify a custom configuration file, any configuration file created in the config directory must be a yaml file and should follow these naming conventions: `config.<custom-naming-of-your-choice>.yml` those additional config files created will not be commited to the git repository.\n").
	PlainText("To start the app you can do `make run` which translates to the following:").
	CodeBlocks("shell", `if ! [ -f ./config/config.dev.yml ];    \
then \
	go run ./main.go;    \
else \
	go run ./main.go --config-file=./config/config.dev.yml;    \
fi`).
	PlainText("\nSo if you have a `config.dev.yml` file it will use it, if not, it will use the default one. To specify another config file which isn't a `config.dev.yml` file run:").
	CodeBlocks("shell", "go run ./main.go --config-file=<name-of-your-config-file>.yml").
	PlainText("\nNote that your custom config file should be placed in the config directory, if it's not placed there, this command will fail.\n").
	H3("Watch Functions\n").
	PlainText("If you're in development flow and wants the app to restart everytime you save a file, follow these steps:").
	OrderedList(
		"Install [GoW](https://github.com/mitranim/gow)",
		"Run `make watch` (it does the same checks as `make run` to use the `config.dev.yml`)",
		"To stop, use `ctrl + c`",
	).
	PlainText("\nTo run with a custom config file and watch functions:\n").
	CodeBlocks("shell", "gow -c -v -r=false run ./main.go --config-file=<name-of-your-config-file>.yml").
	PlainText("\nNote that your custom config file should be placed in the config directory, if it's not placed there, this command will fail.\n").
	H2("Utilities\n").
	PlainText("Common commands are contained in the `Makefile`. If it's ever needed, the variables below can be overriden using environment variables:\n").
	CodeBlocks("shell", `GOOSE_MIGRATION_DIR (default: './database/migrations')
GOOSE_DRIVER        (default: 'postgres')
GOOSE_DBSTRING      (default: 'host=localhost port=5432 user=postgres password=postgres dbname=templateapi sslmode=disable')`).
	PlainText("\nFrequently used commands:\n").
	H3("Create a migration").
	CodeBlocks("shell", "make migration NAME='name_of_the_migration'").
	H3("\nMigrate DB to latest version").
	CodeBlocks("shell", "make migrations_up").
	H3("\nRollback database by one migration").
	CodeBlocks("shell", "make migrations_down").
	H3("\nReset the database").
	CodeBlocks("shell", "make reset_db").
	PlainText("\nCheck the `Makefile` for more commands").
	String()

var READMETemplate = template.Must(template.New("README.md.tmpl").Parse(READMEContent))
