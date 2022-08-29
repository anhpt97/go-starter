package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDb),
	fx.Provide(NewEnv),
)
