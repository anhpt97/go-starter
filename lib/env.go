package lib

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Env struct {
	PORT int

	MONGODB_URI string
	DB_NAME     string

	JWT_EXPIRES_AT time.Duration
	JWT_SECRET     []byte
}

func NewEnv(configPaths ...string) (env Env) {
	configPath := ".env"
	if len(configPaths) > 0 {
		configPath = configPaths[0]
	}
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	viper.AutomaticEnv()

	viper.SetDefault("PORT", 3000)

	viper.SetDefault("MONGODB_URI", "mongodb+srv://root:1@test.dnhgygq.mongodb.net")
	viper.SetDefault("DB_NAME", "test")

	viper.SetDefault("JWT_EXPIRES_AT", 86400)
	viper.SetDefault("JWT_SECRET", []byte("ptanh97"))

	viper.Unmarshal(&env, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc()))
	return
}
