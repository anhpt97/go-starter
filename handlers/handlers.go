package handlers

import "github.com/google/wire"

var Set = wire.NewSet(
	NewAuthHandler,
	NewBookHandler,
)
