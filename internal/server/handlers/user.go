package handlers

import (
	"github.com/gin-gonic/gin"

	"RESTful-Api-Golang/internal/config/database"
	"RESTful-Api-Golang/pkg/controllers"
)

type UserHandler struct {
	db database.Service
}

func NewUserHandler(db database.Service) *UserHandler {
	return &UserHandler{db: db}
}

func (s *UserHandler) UserConnected(c *gin.Context) {
	userData, _, err := controllers.GetUserDataFromContext(c, "user")
	if err != nil {
		controllers.HttpResponse(c, 500, err.Error(), nil)
		return
	}

	var resp string = "Hello, " + userData.Username + " | uuid: " + userData.Uuid

	controllers.HttpResponse(c, 200, resp, nil)
}
