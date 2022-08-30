package lib

import (
	"fmt"
	"go-starter/entities"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db struct {
	*gorm.DB
}

func NewDb(env Env) Db {
	var (
		username = env.DB_USER
		password = env.DB_PASS
		host     = env.DB_HOST
		port     = env.DB_PORT
		dbname   = env.DB_NAME
	)

	dsn := username + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbname + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true"
	fmt.Println(dsn)

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	entities.Sync(db)

	return Db{db}
}
