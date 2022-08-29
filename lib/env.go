package lib

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Env struct {
	PORT int

	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT int
	DB_NAME string

	JWT_EXPIRES_AT time.Duration
	JWT_SECRET     []byte
}

func NewEnv() (env *Env) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.AutomaticEnv()

	viper.SetDefault("PORT", 3000)

	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASS", "")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 3306)
	viper.SetDefault("DB_NAME", "test")

	viper.SetDefault("JWT_EXPIRES_AT", 86400)
	viper.SetDefault("JWT_SECRET", []byte("ptanh97"))

	viper.Unmarshal(&env, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc()))
	return
}
