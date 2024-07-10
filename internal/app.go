package internal

import (
	"application/internal/domain/user"
	"application/internal/infra"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var App = fx.New(
	fx.Provide(infra.NewServer),
	fx.Provide(infra.NewDatabase),
	fx.Invoke(func(db *gorm.DB) {}),
	fx.Invoke(func(server *infra.Server) {}),
	user.UsersModule,
)
