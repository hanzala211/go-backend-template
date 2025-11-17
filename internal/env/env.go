package env

import "os"

func GetEnv(key string, fallback string) string {
	str := os.Getenv(key)
	if str == "" {
		return fallback
	}
	return str
}
