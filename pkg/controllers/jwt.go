package controllers

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"RESTful-Api-Golang/internal/config"
	"RESTful-Api-Golang/internal/models"
)

var secretStr, _ = config.LoadEnv("API_SECRET_KEY")
var secretKey = []byte(secretStr)

func CreateLoginToken(uuid string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"uuid": uuid,
			"exp":  jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // 1 hour
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckLoginTokenData(tokenString string) (string, error) {

	str, token := ParseTokenErrorChecking(tokenString)
	if token == nil {
		return "", errors.New(str)
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	uuid, _ := claims["uuid"].(string)

	return uuid, nil
}

func GetTokenData(tokenString string) (models.TokenDataStruct, error) {

	str, token := ParseTokenErrorChecking(tokenString)
	if token == nil {
		return models.TokenDataStruct{}, errors.New(str)
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	tokenData := models.TokenDataStruct{
		Token: tokenString,
		Uuid: claims["uuid"].(string),
		Exp:  claims["exp"].(string),
	}

	return tokenData, nil

}

func ParseTokenErrorChecking(tokenString string) (string, *jwt.Token) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return secretKey, nil
	})

	switch {
	case token.Valid:
		return "", token
	case errors.Is(err, jwt.ErrSignatureInvalid):
		return "Invalid signature", nil
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return "Invalid token signature", nil
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return "Token is expired", nil
	default:
		return "Error parsing token", nil
	}

}

func GetUserDataFromContext(c *gin.Context, context string) (*models.User, *models.TokenDataStruct, error) {

	var userData models.User = models.User{}
	var tokenData models.TokenDataStruct = models.TokenDataStruct{}

	if context == "user" {

		user, ok := c.Get("user")
		if user == nil && !ok {
			return nil, nil, errors.New("User not found in context")
		}

		userData, _ = user.(models.User)

	} else if context == "token" {

		token, ok := c.Get("token")
		if token == nil && !ok {
			return nil, nil, errors.New("Token not found in context")
		}

		tokenData, _ = GetTokenData(token.(string))

	}
	return &userData, &tokenData, nil
}
