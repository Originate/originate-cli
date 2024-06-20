package migrations

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/pressly/goose/v3"
)

func Down(i MigrationsInput) int {
	provider, err := createPostgresProvider(i.MigrationsDir)
	if err != nil {
		fmt.Printf("Failed to create goose provider: %s\n", err.Error())
		return 1
	}

	ctx, cancel := context.WithTimeout(i.Context, time.Second*20)
	defer cancel()

	result, err := provider.Down(ctx)
	if err != nil {
		if errors.Is(err, goose.ErrNoNextVersion) {
			fmt.Println("No migrations to revert!")
			return 0
		}

		fmt.Printf("Failed to apply migrations: %s\n", err.Error())
		return 1
	}

	fmt.Println("The following migration was reverted successfully:")
	fmt.Println(result.Source.Path)

	return 0
}
