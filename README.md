# Originate CLI

This is a CLI made to accelerate the workflow of the Originate team when developing solutions.

## Requirements

- Go 1.22+

## Installing the CLI
```shell
go install github.com/Originate/originate-cli/cmd/originate@latest
```

## Usage
The CLI is currently able to generate Golang APIs according to the Originate template found [here](https://github.com/Originate/go-api-template) and to manage migrations using the Provider feature from goose.

### API Generation
To generate a new API use the following command:
```shell
originate new api --name example-api --org my-company
```

The `--org` flag isn't required and if it's not provided it defaults to **Originate**

The `--name` flag is required and should obey the following:
- The name of the api should always be all lowercase and ending with -api, ex: **user-api**, **payment-api** and etc.

### Migrations Management
The migrations management is achieved through the [goose Provider feature](https://pressly.github.io/goose/blog/2023/goose-provider/)

The API currently supports the following actions on migrations:
- create
- up to latest version
- down by one version
- reset database

The CLI currently only supports `.sql` migrations.

The commands usage are supposed to be the following:

#### Creation
```shell
originate migrations new --name create_user --migrations-dir database/migrations
```

The `--name` flag is required for the migration creation

The `--migrations-dir` is the directory on where the migrations for the project reside, it's supposed to be a relative path but without the `./` part, it's an optional flag and the default value is `database/migrations`. The directory must exist before creating the migration, or else the command will fail.

#### Up
```shell
originate migrations up --migrations-dir database/migrations
```

The `--migrations-dir` flag is the directory on where the migrations for the project reside, it's supposed to be a relative path but without the `./` part, it's an optional flag and the default value is `database/migrations`. The directory must exist before creating the migration, or else the command will fail.

#### Down
```shell
originate migrations down --migrations-dir database/migrations
```

The `--migrations-dir` flag is the directory on where the migrations for the project reside, it's supposed to be a relative path but without the `./` part, it's an optional flag and the default value is `database/migrations`. The directory must exist before creating the migration, or else the command will fail.

#### Reset
```shell
originate migrations reset --migrations-dir database/migrations
```

The `--migrations-dir` flag is the directory on where the migrations for the project reside, it's supposed to be a relative path but without the `./` part, it's an optional flag and the default value is `database/migrations`. The directory must exist before creating the migration, or else the command will fail.
