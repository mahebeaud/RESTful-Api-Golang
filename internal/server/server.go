package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"RESTful-Api-Golang/internal/config"
	"RESTful-Api-Golang/internal/config/database"
	"RESTful-Api-Golang/internal/models"
)

type Server struct {
	port int
	Db   database.Service
}


var conf = config.GetConfig()

func NewServer() *http.Server {
	port, _ := strconv.Atoi(conf.Port)
	NewServer := &Server{
		port: port,
		Db:   database.ConnectDB(),
	}

	// Init migration after db connection
	NewServer.Db.MakeMigration(&models.User{})

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
