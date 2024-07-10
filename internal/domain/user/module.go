package user

import "go.uber.org/fx"

var UsersModule = fx.Module("users",
	fx.Provide(
		NewUserRouter,
		NewUserService,
	),
	fx.Invoke(NewUserRouter),
)
