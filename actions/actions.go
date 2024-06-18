package actions

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/Originate/originate-cli/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var templateFiles = map[string]string{
	"go.mod.tmpl":  "go.mod",
	"main.go.tmpl": "main.go",
	"README.md":    "README.md",
}

var substitutions = map[string]string{
	"<module>":               "Name",
	"<github-organization>":  "Organization",
	"Originate Template API": "Name",
}

type GenerateAPIConfig struct {
	Name         string
	Organization string
	GithubUser   string
	GithubToken  string
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

	tmpDir, err := os.MkdirTemp(os.TempDir(), "")
	if err != nil {
		fmt.Printf("failed to create tmp dir %s", err)
		return 1
	}

	_, err = git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL:           "https://github.com/Originate/go-api-template.git",
		Depth:         1,
		SingleBranch:  true,
		ReferenceName: plumbing.Master,
	})
	if err != nil {
		if err.Error() == "authentication required" {
			fmt.Println("Check the README on how to set up the authentication step")
		} else {
			fmt.Printf("failed to clone template repository %s\n", err)
		}

		return 1
	}

	if err := utils.CopyDir(tmpDir, config.Name); err != nil {
		fmt.Printf("error copying files to project directory %s\n", err)
		return 1
	}

	for templateName, newName := range templateFiles {
		finalName := fmt.Sprintf("%s/%s", config.Name, newName)
		currentName := fmt.Sprintf("%s/%s", config.Name, templateName)

		if err := os.Rename(currentName, finalName); err != nil {
			fmt.Printf("error renaming %s: %s", templateName, err)
			return 1
		}

		var finalContent string
		contents, err := os.ReadFile(finalName)
		if err != nil {
			fmt.Printf("error reading %s file: %s", templateName, err)
			return 1
		}

		finalContent = string(contents)

		for toSubstitue, fieldName := range substitutions {
			substitutionValue := reflect.ValueOf(config).FieldByName(fieldName).String()

			finalContent = strings.Replace(finalContent, toSubstitue, substitutionValue, -1)
		}

		os.WriteFile(finalName, []byte(finalContent), 0755)
	}

	if err := os.RemoveAll(tmpDir); err != nil {
		fmt.Printf("project set up but the tmp dir wasn't removed please remove it manually: %s", tmpDir)
		return 1
	}

	fmt.Printf("Project %s succesfully set up!\n", config.Name)

	return 0
}
