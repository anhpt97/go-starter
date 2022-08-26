package main

import (
	"go-starter/database"
	_ "go-starter/docs"
	"go-starter/env"
	"go-starter/handlers"
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
		repositories.Module,
		fx.Invoke(http.ListenAndServe(":"+env.PORT, r))).Run()
	// log.Fatal(http.ListenAndServe(":"+env.PORT, r))
}
