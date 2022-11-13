package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	APIPort     string
	APIKey      string
	TokenSecret string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	Cfg = cfg

}

/*
var (
	APIPort     = SetEnv("APIPort", ":8080")
	APIKey      = SetEnv("APIKey", "UwawPangkat2")
	TokenSecret = "ApaIhLiatLiat"
)

func SetEnv(key, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}
*/
