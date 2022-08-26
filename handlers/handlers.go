package handlers

import (
	"go-starter/repositories"

	"go.uber.org/fx"
)

var (
	bookRepository = repositories.BookRepository{}
	userRepository = repositories.UserRepository{}
)

var Module = fx.Options(
	fx.Provide(NewAuthHandler),
)
