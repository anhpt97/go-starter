package handlers

import "go.uber.org/fx"

var Module = fx.Provide(
	NewAuthHandler,
	NewBookHandler,
	NewFileHandler,
	NewHealthHandler,
	NewMeHandler,
	NewUserHandler,
)
