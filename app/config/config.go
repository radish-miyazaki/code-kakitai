package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var (
	once   sync.Once
	config Config
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
	ReadDB ReadDBConfig
	Redis  RedisConfig
}

type DBConfig struct {
	Name     string `envconfig:"DB_NAME" default:"code_kakitai"`
	User     string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_PASSWORD" default:""`
	Host     string `envconfig:"DB_HOST" default:"db"`
	Port     string `envconfig:"DB_PORT" default:"3306"`
}

type ReadDBConfig struct {
	Name     string `envconfig:"READ_DB_NAME" default:"code_kakitai"`
	User     string `envconfig:"READ_DB_USER" default:"root"`
	Password string `envconfig:"READ_DB_PASSWORD" default:""`
	Host     string `envconfig:"READ_DB_HOST" default:"db"`
	Port     string `envconfig:"READ_DB_PORT" default:"3306"`
}

type RedisConfig struct {
	Host string `envconfig:"REDIS_HOST" default:"redis"`
	Port string `envconfig:"REDIS_PORT" default:"6379"`
}

type ServerConfig struct {
	Address string `envconfig:"ADDRESS" default:"0.0.0.0"`
	Port    string `envconfig:"PORT" default:"8000"`
}

func GetConfig() *Config {
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			log.Panic(err)
		}
	})

	return &config
}
