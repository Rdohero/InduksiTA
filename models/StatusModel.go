package models

type Status struct {
	StatusID       uint   `gorm:"primarykey" json:"service_report_id"`
	StatusName     string `json:"status_name"`
	ServiceReports []ServiceReports
}
