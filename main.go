package main

import (
	"go-starter/database"
	_ "go-starter/docs"
	"go-starter/env"
	"go-starter/handlers"
	"go-starter/lib"
	"go-starter/repositories"
	"go-starter/routers"
	"net/http"

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
	r := routers.New()
	fx.New(
		handlers.Module,
		lib.Module,
		repositories.Module,
		fx.Invoke(http.ListenAndServe(":"+env.PORT, r)),
	).Run()
}
