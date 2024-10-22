package handlers

import (
	"log"

	"github.com/gin-gonic/gin"

	"RESTful-Api-Golang/internal/config/database"
	"RESTful-Api-Golang/internal/models"
	"RESTful-Api-Golang/pkg/controllers"
)

type AuthHandler struct {
	db database.Service
}

func NewAuthHandler(db database.Service) *AuthHandler {
	return &AuthHandler{db: db}
}

func (s *AuthHandler) LoginHandler(c *gin.Context) {

	var input models.AuthLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		controllers.HttpResponse(c, 400, "Missing body", nil)
		return
	}

	log.Printf("Login attempt with username: %s", input.Username)

	userData, err := s.db.GetUserByUsernameAndPassword(input.Username, input.Password)
	if err != nil {
		controllers.HttpResponse(c, 401, "Invalid username or password", nil)
		return
	}

	token, err := controllers.CreateLoginToken(userData.Uuid)
	if err != nil {
		controllers.HttpResponse(c, 500, "Error while creating token", nil)
		return
	}

	err = s.db.StoreUserLoginToken(token, userData.Uuid)
	if err != nil {
		controllers.HttpResponse(c, 500, "Error while storing token", nil)
		return
	}

	data := map[string]string{
		"token": token,
	}

	controllers.HttpResponse(c, 200, "Login successful", data)
}

func (s *AuthHandler) RegisterHandler(c *gin.Context) {

	var input models.AuthRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		controllers.HttpResponse(c, 400, "Missing body", nil)
		return
	}

	if input.Password != input.ConfirmPassword {
		controllers.HttpResponse(c, 400, "Passwords does not match", nil)
		return
	}

	err := s.db.IsUsernameTaken(input.Username)
	if err != nil {
		controllers.HttpResponse(c, 401, err.Error(), nil)
		return
	}

	if err := s.db.IsEmailTaken(input.Email); err != nil {
		controllers.HttpResponse(c, 401, err.Error(), nil)
		return
	}

	hashedPassword, _ := controllers.HashPassword(input.Password)
	userUuid := controllers.GenerateUuid("user")

	newUser := &models.User{
		Uuid:      userUuid,
		Email:     input.Email,
		Username:  input.Username,
		Password:  hashedPassword,
		CreatedAt: controllers.GetCurrentDateISO8601(),
	}

	if err := s.db.CreateUser(newUser); err != nil {
		controllers.HttpResponse(c, 400, "Error while creating user", nil)
		return
	}

	controllers.HttpResponse(c, 200, "User created - You can now login", nil)
}

func (s *AuthHandler) LogoutHandler(c *gin.Context) {

	var input models.AuthLogout
	err := c.ShouldBindJSON(&input)
	if err != nil {
		controllers.HttpResponse(c, 400, "Missing body", nil)
		return
	}

	_, dataToken, err := controllers.GetUserDataFromContext(c, "token")
	if err != nil {
		controllers.HttpResponse(c, 500, "Error while getting token", nil)
		return
	}

	if input.Uuid != dataToken.Uuid {
		controllers.HttpResponse(c, 401, "Invalid user", nil)
		return
	}

	err = s.db.CheckIfRightUserLogout(input.Uuid, dataToken.Token)
	if err != nil {
		controllers.HttpResponse(c, 401, "Already disconnected", nil)
		return
	}

	err = s.db.DeleteUserLoginToken(input.Uuid, dataToken.Token)
	if err != nil {
		controllers.HttpResponse(c, 500, "Error while deleting token", nil)
		return
	}

	controllers.HttpResponse(c, 200, "Logout successful", nil)

}
