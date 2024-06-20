package migrations

import (
	"context"
	"fmt"
	"time"
)

type MigrationsInput struct {
	MigrationsDir string
	Context       context.Context
}

func Up(i MigrationsInput) int {
	provider, err := createPostgresProvider(i.MigrationsDir)
	if err != nil {
		fmt.Printf("Failed to create goose provider: %s\n", err.Error())
		return 1
	}

	ctx, cancel := context.WithTimeout(i.Context, time.Second*20)
	defer cancel()

	hasPending, err := provider.HasPending(ctx)
	if err != nil {
		fmt.Printf("Failed to check migrations on database: %s\n", err.Error())
		return 1
	}

	if !hasPending {
		fmt.Println("All migrations are up to date!")
		return 0
	}

	results, err := provider.Up(ctx)
	if err != nil {
		fmt.Printf("Failed to apply migrations: %s\n", err.Error())
		return 1
	}

	fmt.Println("The following migrations were executed successfully:")
	for _, result := range results {
		fmt.Println(result.Source.Path)
	}

	return 0
}
