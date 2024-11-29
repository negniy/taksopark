package DTO

import (
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

type CarWithModel struct {
	CarID        uint   `gorm:"column:car_id" json:"car_id"`
	LicensePlate string `gorm:"column:license_plate" json:"license_plate"`
	ModelName    string `gorm:"column:model_name" json:"model_name"`
	Year         uint   `gorm:"column:year" json:"year"`
	Manufacturer string `gorm:"column:manufacturer" json:"manufacturer"`
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
	Min int     `json:"min" gorm:"column:min"`
	Max int     `json:"max" gorm:"column:max"`
	Avg float32 `json:"avg" gorm:"column:avg"`
}
