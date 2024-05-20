package models

import "time"

type SalesReports struct {
	SalesReportID uint      `gorm:"primarykey" json:"sales_report_id"`
	Date          time.Time `json:"date"`
	ItemName      string    `json:"item_name"`
	Quantity      int       `json:"quantity"`
}
