package actions

import (
	"fmt"
	"os"
	"regexp"

	"github.com/Originate/originate-cli/config"
	"github.com/Originate/originate-cli/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func GenerateNewAPI(name string, cfg config.Config) int {
	if name == "" {
		fmt.Println("the project needs a name, ex: example-api")
		return 1
	}

	nameRegexValidation := regexp.MustCompile("^[a-z]+(-api){1}$")
	if !nameRegexValidation.Match([]byte(name)) {
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

	if err := utils.CopyDir(tmpDir, name); err != nil {
		fmt.Printf("error copying files to project directory %s\n", err)
		return 1
	}

	if err := os.RemoveAll(tmpDir); err != nil {
		fmt.Printf("project set up but the tmp dir wasn't removed please remove it manually: %s", tmpDir)
		return 1
	}

	fmt.Printf("Project %s succesfully set up!\n", name)

	return 0
}
