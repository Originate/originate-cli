# Originate CLI

This is a CLI made to generate projects using the repository templates from the Originate organization.

## Requirements

- Go 1.22+

Install it with:
```
go install github.com/Originate/originate-cli/cmd/originate@latest
```

## Setup

To test it just run:

```
go run ./cmd/originate/main.go new --name example-api
```

It will create a example-api app using the go-api-template repository from Originate.