package config

import "os"

type Application struct {
	ProductionType  string
	SigningKey      string
	TokenExpiration string
}

type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Redis struct {
	Port     string
	Password string
}

type Config struct {
	Application Application
	Db          Db
	Redis       Redis
}

func NewConfig() *Config {
	return &Config{
		Application: Application{
			ProductionType:  os.Getenv("PRODUCTION_TYPE"),
			SigningKey:      os.Getenv("SIGNING_KEY"),
			TokenExpiration: os.Getenv("TOKEN_EXPIRATION"),
		},
		Db: Db{
			Host:     os.Getenv("POSTGRESQL_HOST"),
			Port:     os.Getenv("POSTGRESQL_PORT"),
			User:     os.Getenv("POSTGRESQL_USER"),
			Password: os.Getenv("POSTGRESQL_PASSWORD"),
			Database: os.Getenv("POSTGRESQL_DB"),
		},
		Redis: Redis{
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}
}
