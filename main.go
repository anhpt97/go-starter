package main

import (
	"go-starter/database"
	_ "go-starter/docs"
	"go-starter/env"
	"go-starter/routers"
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe(":"+env.PORT, r))
}
