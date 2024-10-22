package controllers

import (
	"strings"
	"github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

func HashPassword(password string) (string, error) {

    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func ComparePasswordHash(password string, hash string) bool {

    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
        return false
    }
    return true
}

func GenerateUuid(prefix string) string {

	baseUuid := uuid.NewString()
    uuid := fmt.Sprintf("%s_%s", prefix, strings.ReplaceAll(baseUuid, "-", ""))
    return uuid
}