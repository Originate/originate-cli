package api

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/Originate/originate-cli/actions/api/templates"
	"github.com/Originate/originate-cli/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var templateFiles = map[string]*template.Template{
	"go.mod":    templates.GoModTemplate,
	"main.go":   templates.MainGoTemplate,
	"README.md": templates.READMETemplate,
}

type GenerateAPIConfig struct {
	Name         string
	Organization string
}

type TemplateVariables struct {
	Module        string
	FormattedName string
	Organization  string
}

func GenerateNewAPI(config GenerateAPIConfig) int {
	if config.Name == "" {
		fmt.Println("the project needs a name, ex: example-api")
		return 1
	}

	nameRegexValidation := regexp.MustCompile("^[a-z]+(-api){1}$")
	if !nameRegexValidation.Match([]byte(config.Name)) {
		fmt.Println("name should obey the following parameters: all lowercase and ending in -api")
		return 1
	}

	fmt.Println("Cloning template repository...")

	tmpDir, err := os.MkdirTemp(os.TempDir(), "")
	if err != nil {
		fmt.Printf("failed to create tmp dir %s\n", err)
		return 1
	}

	_, err = git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL:           "https://github.com/Originate/go-api-template.git",
		Depth:         1,
		SingleBranch:  true,
		ReferenceName: plumbing.Master,
	})
	if err != nil {
		fmt.Printf("Failed to clone template repository %s\n", err)
		return 1
	}

	fmt.Println("Setting up base files...")

	if err := utils.CopyDir(tmpDir, config.Name); err != nil {
		fmt.Printf("Failed to copy files to project directory %s\n", err)
		return 1
	}

	vars := TemplateVariables{
		Organization:  config.Organization,
		Module:        config.Name,
		FormattedName: utils.FormatNameForREADME(config.Name),
	}
	for filename, template := range templateFiles {
		path := filepath.Join(config.Name, filename)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			fmt.Printf("Failed to stat path: %s\nerr: %s\n", path, err.Error())
			return 1
		}

		f, err := os.Create(path)
		if err != nil {
			fmt.Printf("Failed to create file %s error: %s\n", filename, err.Error())
			return 1
		}
		defer f.Close()

		if err := template.Execute(f, vars); err != nil {
			fmt.Printf("Failed to template file %s error: %s\n", filename, err.Error())
			return 1
		}
	}

	if err := os.RemoveAll(tmpDir); err != nil {
		fmt.Printf("Project set up but the tmp dir wasn't removed please remove it manually: %s\n", tmpDir)
		return 1
	}

	fmt.Println("Updating go.mod dependencies...")

	if err := os.Chdir(config.Name); err != nil {
		fmt.Printf("Failed to set up go.mod update, please run \"go get .\" manually, error: %s\n", err.Error())
		return 1
	}

	cmd := exec.Command("go", "get", ".")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to update go.mod, please run \"go get .\" manually, error: %s\n", err.Error())
		return 1
	}

	os.Chdir("..")

	fmt.Printf("Project %s succesfully set up!\n", config.Name)

	return 0
}
