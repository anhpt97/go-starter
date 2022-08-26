package handlers

import (
	"go-starter/repositories"

	"github.com/google/wire"
)

var (
	bookRepository = repositories.BookRepository{}
	userRepository = repositories.UserRepository{}
)

var Set = wire.NewSet(
	NewAuthHandler,
)
