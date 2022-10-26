package main

import (
	"context"
	_ "go-starter/docs"
	"go-starter/handlers"
	"go-starter/i18n"
	"go-starter/lib"
	"go-starter/middlewares"
	"go-starter/repositories"
	"go-starter/routers"
	"go-starter/swagger"
	"log"
	"net/http"
	"strconv"

	"go.uber.org/fx"
)

// @Title       Go starter
// @Version     1.0
// @Description Go starter's API documentation

// @SecurityDefinitions.apiKey Bearer
// @In                         header
// @Name                       Authorization
func main() {
	fx.New(
		handlers.Module,
		i18n.Module,
		lib.Module,
		middlewares.Module,
		repositories.Module,
		routers.Module,
		fx.Invoke(
			func(lc fx.Lifecycle, routers routers.Routers, env lib.Env) {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) (err error) {
						go func() {
							r := routers.New("/api")
							swagger.New(r, "/swagger")
							log.Fatal(http.ListenAndServe(":"+strconv.Itoa(env.PORT), r))
						}()
						return
					},
				})
			},
		),
	).Run()
}
