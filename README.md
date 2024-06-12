# Originate CLI

This is a CLI made to generate projects using the repository templates from the Originate organization.

## Requirements

- Go 1.22+

Install it with:
```
go install github.com/Originate/originate-cli/cmd/originate@latest
```

## Setup

To set it up you need to set two environment variables:

- ORIGINATE_CLI_GITHUB_USERNAME: your Github username, noting that you need to be a member of the Originate organization to use the CLI
- ORIGINATE_CLI_GITHUB_TOKEN: a authentication token needed to clone the template repos

To test it just run:

```
go run ./cmd/originate/main.go new --name example-api
```

It will create a example-api app using the go-api-template repository from Originate.