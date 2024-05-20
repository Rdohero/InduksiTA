package models

type User struct {
	UserID      uint   `gorm:"primarykey" json:"user_id"`
	Image       string `form:"image" json:"image"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	NoHandphone string `json:"no_handphone"`
}
