package repositories

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewBookRepository),
	fx.Provide(NewUserRepository),
)
