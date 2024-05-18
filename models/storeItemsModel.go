package models

type StoreItems struct {
	StoreItemsID      uint   `gorm:"primarykey" json:"store_items_id"`
	StoreItemsName    string `json:"store_items_name"`
	Quantity          int    `json:"quantity"`
	Price             int    `json:"price"`
	CategoryMachineID uint   `json:"category_machine_id"`
	CategoryMachine   CategoryMachine
}
