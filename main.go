package main

import (
	"context"

	"github.com/Santiago1010/inventory-go/database"
	"github.com/Santiago1010/inventory-go/internal/service"
	"github.com/Santiago1010/inventory-go/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
      service.New,
		),
		fx.Invoke(),
	)

	app.Run()
}
