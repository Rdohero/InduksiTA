package models

import "time"

type ServiceReports struct {
	ServiceReportID     uint                  `gorm:"primarykey" json:"service_report_id"`
	Date                time.Time             `json:"date"`
	Image               string                `form:"image" json:"image"`
	Name                string                `json:"name"`
	MachineName         string                `json:"machine_name"`
	Complaints          string                `json:"complaints"`
	TotalPrice          int                   `json:"total_price"`
	DateEnd             *time.Time            `json:"date_end"`
	StatusID            uint                  `json:"status_id"`
	UserID              uint                  `json:"user_id"`
	Status              Status                `gorm:"foreignKey:StatusID;references:StatusID"`
	User                User                  `gorm:"foreignKey:UserID;references:UserID"`
	ServiceReportsItems []ServiceReportsItems `gorm:"foreignKey:ServiceReportID;references:ServiceReportID"`
}

type ServiceReportsItems struct {
	ServiceReportsItemsID uint           `gorm:"primarykey" json:"service_reports_items_id"`
	StoreItemsID          uint           `json:"store_items_id"`
	ItemName              string         `json:"item_name"`
	Quantity              int            `json:"quantity"`
	Price                 int            `json:"price"`
	Category              string         `json:"category"`
	CategoryID            uint           `json:"category_id"`
	ServiceReportID       uint           `json:"service_report_id"`
	ServiceReports        ServiceReports `gorm:"foreignKey:ServiceReportID;references:ServiceReportID"`
	Categories            Category       `gorm:"foreignKey:CategoryID;references:CategoryID"`
}
