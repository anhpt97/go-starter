package lib

import (
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	PORT string

	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string

	JWT_EXPIRES_AT time.Duration
	JWT_SECRET     []byte
}

func NewEnv() (env Env) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.SetDefault("PORT", 3000)

	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASS", "")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_NAME", "test")

	viper.AutomaticEnv()

	env.JWT_EXPIRES_AT = viper.GetDuration("JWT_EXPIRES_AT")
	env.JWT_SECRET = []byte(viper.GetString("JWT_SECRET"))

	viper.Unmarshal(&env)
	return
}
