package repositories

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var DB *gorm.DB

func New(db *gorm.DB) {
	DB = db
}

func CreateSqlBuilder(model any) *gorm.DB {
	return DB.Model(model)
}

var Set = wire.NewSet(
	NewBookRepository,
	NewUserRepository,
)
