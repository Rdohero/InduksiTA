package models

import "time"

type ServiceReports struct {
	ServiceReportID uint      `gorm:"primarykey" json:"service_report_id"`
	Date            time.Time `json:"date"`
	PersonName      string    `json:"person_name"`
	MachineNumber   string    `json:"machine_number"`
	MachineName     string    `json:"machine_name"`
	Complaints      string    `json:"complaints"`
	Status          string    `json:"status"`
}
