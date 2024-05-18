package models

type CategoryMachine struct {
	CategoryMachineID   uint         `gorm:"primarykey" json:"category_machine_id"`
	CategoryMachineName string       `json:"category_machine_name"`
	StoreItems          []StoreItems `gorm:"foreignKey:CategoryMachineID;references:CategoryMachineID"`
}
