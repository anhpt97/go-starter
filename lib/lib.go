package lib

import "go.uber.org/fx"

var Module = fx.Provide(
	NewDb,
	NewEnv,
)
