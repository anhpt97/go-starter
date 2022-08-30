package main

import (
	"context"
	_ "go-starter/docs"
	"go-starter/handlers"
	"go-starter/lib"
	"go-starter/middlewares"
	"go-starter/repositories"
	"go-starter/routers"
	"net/http"
	"strconv"

	"go.uber.org/fx"
)

// @title       Go starter
// @version     1.0
// @description Go starter's API documentation

// @securityDefinitions.apikey Bearer
// @in   header
// @name Authorization
func main() {
	fx.New(
		handlers.Module,
		lib.Module,
		middlewares.Module,
		repositories.Module,
		routers.Module,
		fx.Invoke(
			func(lc fx.Lifecycle, routers routers.Routers, env lib.Env) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) (err error) {
						go func() {
							r := routers.New("/api/v1")
							http.ListenAndServe(":"+strconv.Itoa(env.PORT), r)
						}()
						return
					},
				})
			},
		),
	).Run()
}
