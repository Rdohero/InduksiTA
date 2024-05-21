package models

type CategoryMachine struct {
	CategoryMachineID   uint         `gorm:"primarykey" json:"category_machine_id"`
	CategoryMachineName string       `json:"category_machine_name"`
	StoreItems          []StoreItems `gorm:"foreignKey:CategoryMachineID;references:CategoryMachineID"`
}

type CategorySparePart struct {
	CategorySparePartID   uint        `gorm:"primarykey" json:"category_spare_part_id"`
	CategorySparePartName string      `json:"category_spare_part_name"`
	SparePart             []SparePart `gorm:"foreignKey:CategorySparePartID;references:CategorySparePartID"`
}
