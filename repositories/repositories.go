package repositories

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var DB *gorm.DB

func New(db *gorm.DB) {
	DB = db
}

func CreateSqlBuilder(model any) *gorm.DB {
	return DB.Model(model)
}

var Module = fx.Options(
	fx.Provide(NewUserRepository),
)
