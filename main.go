package main

import (
	"context"
	"go-starter/database"
	_ "go-starter/docs"
	"go-starter/handlers"
	"go-starter/lib"
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
	database.Connect()
	fx.New(
		handlers.Module,
		lib.Module,
		repositories.Module,
		fx.Invoke(
			func(lc fx.Lifecycle, env lib.Env) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) (err error) {
						go func() {
							r := routers.New()
							http.ListenAndServe(":"+strconv.Itoa(env.PORT), r)
						}()
						return
					},
				})
			},
		),
	).Run()
}
