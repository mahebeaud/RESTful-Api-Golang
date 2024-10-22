package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"RESTful-Api-Golang/internal/models"
	"RESTful-Api-Golang/pkg/controllers"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Service is an interface that represents a service that interacts with a database.

	// MakeMigration runs the migrations for the provided models.
	// Add / remove / modify the models as needed.
	MakeMigration(models ...interface{})

	// Creates a new user in db.
	CreateUser(user *models.User) error

	// Check if username already taken.
	IsUsernameTaken(username string) error

	// Check if email already taken.
	IsEmailTaken(email string) error

	// Retrieves a user by their username and password.
	// Return user models.
	GetUserByUsernameAndPassword(username string, password string) (*models.User, error)

	// Retrieves a user by their uuid.
	GetUserByUuid(uuid string) (interface{}, error)

	// Store login token in db.
	StoreUserLoginToken(token string, uuid string) error

	// Delete login token from db.
	DeleteUserLoginToken(uuid string, token string) error

	// Check if right user logout
	CheckIfRightUserLogout(uuid string, token string) error
}

type service struct {
	db *gorm.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *service
)

func ConnectDB() Service {
	if dbInstance != nil {
		return dbInstance
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	dbInstance = &service{
		db: db,
	}

	return dbInstance
}

func (s *service) MakeMigration(models ...interface{}) {
	if err := s.db.AutoMigrate(models...); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

func (s *service) CreateUser(user *models.User) error {

	if err := s.db.Model(&models.User{}).Create(user).Error; err != nil {
		//if err := s.db.Where("uuid = ?", user.Uuid).Delete(&user).Error; err != nil {
		//	return err
		//}
		return err
	}

	return nil
}

func (s *service) IsUsernameTaken(username string) error {

	log.Printf("Username: %s", username)

	var user models.User
	if err := s.db.Select("username").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Username not found so OK
			return nil
		}
		// Other type of error
		return err
	}

	// Username Already exists
	return errors.New("Username not available")
}

func (s *service) IsEmailTaken(email string) error {

	var user models.User
	if err := s.db.Select("email").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Email not found so OK
			return nil
		}
		// Other type of error
		return err
	}

	// Email Already exists
	return errors.New("Email not available")
}

func (s *service) GetUserByUsernameAndPassword(username string, password string) (*models.User, error) {
	var user models.User

	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	log.Printf("User: %v", user)
	matched := controllers.ComparePasswordHash(password, user.Password)
	if !matched {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func (s *service) GetUserByUuid(uuid string) (interface{}, error) {

	var user = models.User{Uuid: uuid}

	if err := s.db.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No user found
			return nil, nil
		}
		// Other type of error
		return nil, err
	}

	// User exists
	return user, nil
}

func (s *service) StoreUserLoginToken(token string, uuid string) error {

	var user = models.User{Uuid: uuid}

	if err := s.db.First(&user).Error; err != nil {
		return err
	}

	if err := s.db.Model(&user).Update("login_token", token).Error; err != nil {
		return err
	}

	if err := s.db.Model(&user).Update("updated_at", controllers.GetCurrentDateISO8601()).Error; err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteUserLoginToken(uuid string, token string) error {
	var user = models.User{Uuid: uuid}

	if err := s.db.Where("login_token = ?", token).First(&user).Error; err != nil {
		return err
	}

	if err := s.db.Model(&user).Update("login_token", "").Error; err != nil {
		return err
	}

	if err := s.db.Model(&user).Update("updated_at", controllers.GetCurrentDateISO8601()).Error; err != nil {
		return err
	}

	return nil
}

func (s *service) CheckIfRightUserLogout(uuid string, token string) error {
	var user = models.User{Uuid: uuid}

	if err := s.db.Where("login_token = ?", token).First(&user).Error; err != nil {
		return err
	}

	return nil
}