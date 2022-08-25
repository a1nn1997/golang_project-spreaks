package config

import (
	"github.com/a1nn1997/serverlesscrud1/utils/env"
	"strconv"
)

type Config struct {
	Port        int
	Timeout     int
	Dialect     string
	DatabaseURI string
}

func GetConfig() Config {
	return Config{
		Port:        parseEnvToInt("PORT", "8069"),
		Timeout:     parseEnvToInt("TIMEOUT", "30"),
		Dialect:     env.GetEnv("DIALECT", "sqlite3"),
		DatabaseURI: env.GetEnv("DATABASE_URI", ":memory:"),
	}
}

func parseEnvToInt(envName, defaultValue string) int {
	num, err := strconv.Atoi(env.GetEnv(envName, defaultValue))

	if err != nil {
		return 0
	}

	return num
}
