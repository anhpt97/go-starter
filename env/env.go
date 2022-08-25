package env

import (
	"time"

	"github.com/spf13/viper"
)

var (
	PORT string

	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string

	JWT_EXPIRES_AT time.Duration
	JWT_SECRET     []byte
)

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	PORT = viper.GetString("PORT")

	DB_USER = viper.GetString("DB_USER")
	DB_PASS = viper.GetString("DB_PASS")
	DB_HOST = viper.GetString("DB_HOST")
	DB_PORT = viper.GetString("DB_PORT")
	DB_NAME = viper.GetString("DB_NAME")

	JWT_EXPIRES_AT = time.Duration(viper.GetInt("JWT_EXPIRES_AT"))
	JWT_SECRET = []byte(viper.GetString("JWT_SECRET"))
}
