package entities

import (
	"go-starter/lib"

	"go.uber.org/fx"
)

func Sync(db lib.Db) {
	db.AutoMigrate(
		Book{},
		User{},
	)
}

var Module = fx.Invoke(Sync)
