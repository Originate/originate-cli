package templates

import "text/template"

var MainGoTemplate = template.Must(template.New("main.go.tmpl").Parse(`package main

import (
	"fmt"
	"log/slog"

	"github.com/Originate/go-utilities/configutilities"
	"github.com/Originate/go-utilities/databaseutilities"
	"github.com/Originate/go-utilities/loggerutilities"
	"github.com/Originate/go-utilities/docutilities"
	"github.com/{{.Organization}}/{{.Module}}/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	_ "github.com/{{.Organization}}/{{.Module}}/docs"  
)

func main() {
	var cfg config.Config

	if err := configutilities.Load(&cfg); err != nil {
		slog.Error("Failed to load config", "err", err)
		panic(err)
	}

	loggerutilities.SetupLogger(cfg.Slog)

	slog.Info("Loaded config", "config", configutilities.HideSensitive(cfg))
	db, err := databaseutilities.NewPostgres(context.Background(), cfg.Database)
	if err != nil {
		slog.Error("Failed to init database", "err", err)
		panic(err)
	}
	defer db.Close()

	validate := validator.New(validator.WithRequiredStructEnabled())
	gin.SetMode(cfg.Gin.Mode)
	router := gin.Default()

	docutilities.InitDocs(router)

	if err := router.Run(fmt.Sprintf(":%d", cfg.Gin.Port)); err != nil {
		slog.Error("Error running the service", "err", err)
		panic(err)
	}
}`))
