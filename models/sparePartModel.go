package models

type SparePart struct {
	SparePartID         uint   `gorm:"primarykey" json:"spare_part_id"`
	SparePartName       string `json:"spare_part_name"`
	Quantity            int    `json:"quantity"`
	Price               int    `json:"price"`
	CategorySparePartID uint   `json:"category_spare_part_id"`
	CategorySparePart   CategorySparePart
}
