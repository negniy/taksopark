package DTO

import (
	"taksopark/internal/models"
	"time"
)

type CarWithModel struct {
	CarID        uint            `json:"car_id" gorm:"column:car_id"`
	LicensePlate string          `json:"license_plate" gorm:"column:license_plate"`
	Model        models.CarModel `json:"model" gorm:"column:model_name"`
	Year         time.Time       `json:"year" gorm:"column:year"`
	Manufacturer string          `json:"manufacturer" gorm:"column:manufacturer"`
}

type PersonCount struct {
	Person `json:"person"`
	Count  uint `json:"count" gorm:"column:count"`
}

type Person struct {
	Name    string `json:"name" gorm:"column:first_name"`
	Surname string `json:"surname" gorm:"column:last_name"`
}

type DriverCount struct {
	PersonCount  `json:"person_count"`
	LicensePlate string `json:"license_plate" gorm:"column:license_plate"`
}

type Statistic struct {
	Min time.Time `json:"min" gorm:"column:min"`
	Max time.Time `json:"max" gorm:"column:max"`
	Avg time.Time `json:"avg" gorm:"column:avg"`
}
