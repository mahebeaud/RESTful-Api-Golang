package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"RESTful-Api-Golang/internal/config/database"
	"RESTful-Api-Golang/pkg/controllers"
)

func VerifyTokenMiddleware(db database.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			controllers.HttpResponse(c, http.StatusUnauthorized, controllers.ErrNotAuthorizedMsg, nil)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || authToken[0] != "Bearer-token" {
			controllers.HttpResponse(c, http.StatusUnauthorized, controllers.ErrNotAuthorizedMsg, nil)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		uuid, err := controllers.CheckLoginTokenData(authToken[1])
		if err != nil {
			controllers.HttpResponse(c, http.StatusUnauthorized, err.Error(), nil)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := db.GetUserByUuid(uuid)
		if err != nil || user == nil {
			controllers.HttpResponse(c, http.StatusNotFound, controllers.ErrNotFoundMsg, nil)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func UserConnectedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		value, exists := c.Get("user")
		if value != nil && exists {
			controllers.HttpResponse(c, http.StatusNotFound, controllers.ErrNotFoundMsg, nil)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}

func GetUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		authToken := strings.Split(authHeader, " ")
		if authHeader == "" {
			c.Next()
		} else {
			if len(authToken) == 2 && authToken[0] == "Bearer-token" {
				c.Set("token", authToken[1])
				c.Next()
			}
		}
		c.Next()
	}
}
