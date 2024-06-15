package models

type Role struct {
	RoleID   uint   `gorm:"primarykey" json:"role_id"`
	RoleName string `json:"role_name"`
	User     []User
}
