package configs

import "os"

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
