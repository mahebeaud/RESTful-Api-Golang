package controllers

import (
	"github.com/gin-gonic/gin"
)

var (
	// error msgs
	ErrNotAuthorizedMsg              = "Not authorized"
	ErrNotFoundMsg                   = "Not found"
	ErrInvalidUsernameOrPasswordMsg  = "Invalid username or password"
	ErrUserAlreadyExistsMsg          = "User already exists"
	ErrPasswordsDoNotMatchMsg        = "Passwords do not match"
	ErrInternalServerMsg             = "Internal server error"
	ErrMissingAuthorizationHeaderMsg = "Missing Authorization header"
)

type GoResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func HttpResponse(c *gin.Context, code int, message string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.Status(code)

	if code != 200 {
		c.JSON(code, GoResponse{
			Code:    code,
			Error:   message,
		})
		return
	} else {
		c.JSON(code, GoResponse{
			Code:  code,
			Message: message,
			Data:  data,
		})
		return
	}
}
