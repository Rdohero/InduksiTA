package models

type Category struct {
	CategoryID   uint         `gorm:"primarykey" json:"category_id"`
	CategoryName string       `json:"category_name"`
	SparePart    []SparePart  `gorm:"foreignKey:CategoryID;references:CategoryID"`
	StoreItems   []StoreItems `gorm:"foreignKey:CategoryID;references:CategoryID"`
}
