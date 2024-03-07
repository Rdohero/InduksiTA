package models

type User struct {
	UserID   uint   `gorm:"primarykey" json:"user_id"`
	Email    string `gorm:"unique;not null" json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}
