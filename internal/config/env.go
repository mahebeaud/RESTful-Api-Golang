package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type EnvConfig struct {
	Env  string
	Port string
}

var (
	cfg  *EnvConfig
	once sync.Once
)

func GetConfig() *EnvConfig {

	once.Do(func() {
		cfg = &EnvConfig{}
		cfg.LoadConfig()
	})
	return cfg
}

func LoadEnv(key string) (string, error) {

	err := godotenv.Load(".env")

	if err != nil {
		return "Error loading .env file", err
	}

	return os.Getenv(key), nil
}

func (c *EnvConfig) LoadConfig() {

	env, err := LoadEnv("ENV")
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	if env == "dev" {
		c.Env = env
		log.Print("Debug mode enabled\n")
		c.Port = "8080"
	} else if env == "prod" {
		c.Env = env
		c.Port = "8080"
		gin.SetMode(gin.ReleaseMode)
	}

	log.Print("Config loaded\n")
}
