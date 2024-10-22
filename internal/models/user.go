package models

type User struct {
	Uuid       string `gorm:"primary_key"`
	Username   string `json:"username" gorm:"unique"`
	Password   string `json:"password" gorm:"not null"`
	Email      string `json:"email" gorm:"unique"`
	LoginToken []string `json:"login_token gorm:type:json"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
