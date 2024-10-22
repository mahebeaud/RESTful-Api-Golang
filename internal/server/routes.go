package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"RESTful-Api-Golang/internal/middleware"
	"RESTful-Api-Golang/internal/server/handlers"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "error": "Not found"})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"code": 405, "error": "Method not allowed"})
	})

	r.GET("/", s.HealthCheck)

	app := r.Group("/api", middleware.GetUserToken())
	{
		auth := app.Group("/auth")
		{
			authHandlers := handlers.NewAuthHandler(s.Db)

			// Base url: /api/auth/*
			auth.POST("/login", authHandlers.LoginHandler, middleware.UserConnectedMiddleware())
			auth.POST("/register", authHandlers.RegisterHandler, middleware.UserConnectedMiddleware())
			auth.POST("/logout", authHandlers.LogoutHandler)
		}

		user := app.Group("/user", middleware.VerifyTokenMiddleware(s.Db))
		{
			userHandlers := handlers.NewUserHandler(s.Db)

			// Base url: /api/user/*
			user.GET("/", userHandlers.UserConnected)
		}
	}

	return r
}

func (s *Server) HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, "OK")
}
