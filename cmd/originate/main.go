package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Originate/originate-cli/actions"
	"github.com/Originate/originate-cli/config"
	"github.com/Originate/originate-cli/utils"
)

func main() {
	help := flag.Bool("help", false, "Prints out the help instructions")
	projectName := flag.String("name", "", "Name of the API to be generated")
	filePath := flag.String("config-file", "./config/config.yml", "The config file to be used")

	// This needs to be below all the flag definitions
	utils.ParseFlags()

	var cfg config.Config
	if err := config.Load(&cfg, *filePath); err != nil {
		fmt.Println("failed to load all config, please check out the usage")
		os.Exit(1)
	}

	if *help {
		fmt.Println("Please use the command as: originate new --name \"example-api\"")
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	// Here ATM it's just generating a backend API but the idea is to have all sorts of boilerplates
	// to generate from like: a full monorepo boilerplate with bezel for multi-language monorepos
	// or a pnpm + turbo managed one for JS/TS monorepos
	action := flag.Arg(0)
	switch action {
	case "new":
		os.Exit(actions.GenerateNewAPI(*projectName, cfg))
	default:
		fmt.Printf("unrecognized action: %s", action)
		os.Exit(1)
	}
}
