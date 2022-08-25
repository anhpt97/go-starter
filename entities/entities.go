package entities

import "gorm.io/gorm"

func Sync(db *gorm.DB) {
	db.AutoMigrate(
		Book{},
		User{},
	)
}
