package handlers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthHandler),
	fx.Provide(NewBookHandler),
	fx.Provide(NewFileHandler),
)
