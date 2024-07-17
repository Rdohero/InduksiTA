package models

import "time"

type SalesReports struct {
	SalesReportID    uint               `gorm:"primarykey" json:"sales_report_id"`
	Date             time.Time          `json:"date"`
	TotalPrice       int                `json:"total_price"`
	SalesReportItems []SalesReportItems `gorm:"foreignKey:SalesReportID;references:SalesReportID"`
}

type SalesReportItems struct {
	SalesReportItemsID uint         `gorm:"primarykey" json:"sales_report_items_id"`
	StoreItemsID       uint         `json:"store_items_id"`
	ItemName           string       `json:"item_name"`
	Quantity           int          `json:"quantity"`
	Price              int          `json:"price"`
	Category           string       `json:"category"`
	CategoryID         uint         `json:"category_id"`
	SalesReportID      uint         `json:"sales_report_id"`
	SalesReports       SalesReports `gorm:"foreignKey:SalesReportID;references:SalesReportID"`
	Categories         Category     `gorm:"foreignKey:CategoryID;references:CategoryID"`
}
