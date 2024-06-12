package config

import (
	"errors"
	"io/fs"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GitHubUser  string `env:"ORIGINATE_CLI_GITHUB_USERNAME" validate:"required" yaml:"username"`
	GitHubToken string `env:"ORIGINATE_CLI_GITHUB_TOKEN"    validate:"required" yaml:"token"`
}

const defaultFilePath = "./config/config.yml"

func Load(cfg interface{}, filePath string) error {
	if err := cleanenv.ReadConfig(filePath, cfg); err != nil {
		if filePath == defaultFilePath && errors.Is(err, fs.ErrNotExist) {
			slog.Debug("Skipping missing config file", "error", err)

			if err := cleanenv.ReadEnv(cfg); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(cfg)
}
