package migrations

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/pressly/goose/v3"
)

func Reset(i MigrationsInput) int {
	provider, err := createPostgresProvider(i.MigrationsDir)
	if err != nil {
		fmt.Printf("Failed to create goose provider: %s\n", err.Error())
		return 1
	}

	ctx, cancel := context.WithTimeout(i.Context, time.Second*20)
	defer cancel()

	for true {
		_, err := provider.Down(ctx)
		if err != nil {
			if errors.Is(err, goose.ErrNoNextVersion) {
				break
			}

			fmt.Printf("Failed to revert migration: %s\n", err.Error())
			return 1
		}
	}

	fmt.Println("Database was successfully reset!")

	return 0
}
