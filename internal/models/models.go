package models

import (
	"time"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

type Driver struct {
	DriverID      uint   `gorm:"primaryKey;autoIncrement" json:"driver_id"`
	FirstName     string `gorm:"size:100" json:"first_name"`
	LastName      string `gorm:"size:100" json:"last_name"`
	LicenseNumber string `gorm:"uniqueIndex" json:"license_number"`
}

type CarModel struct {
	ModelID      uint   `gorm:"primaryKey;autoIncrement" json:"model_id"`
	ModelName    string `gorm:"size:100" json:"model_name"`
	Manufacturer string `gorm:"size:100" json:"manufacturer"`
}

type Car struct {
	CarID        uint     `gorm:"primaryKey;autoIncrement" json:"car_id"`
	LicensePlate string   `gorm:"size:100;uniqueIndex" json:"license_plate"`
	ModelID      uint     `json:"model_id"`
	Model        CarModel `gorm:"foreignKey:ModelID" json:"model"`
	Year         uint     `gorm:"type:year" json:"year"`
	Notes        string   `gorm:"type:json" json:"notes"`
}

type Customer struct {
	CustomerID uint   `gorm:"primaryKey;autoIncrement" json:"customer_id"`
	FirstName  string `gorm:"size:100" json:"first_name"`
	LastName   string `gorm:"size:100" json:"last_name"`
	Phone      string `gorm:"size:15" json:"phone"`
}

type Trip struct {
	TripID     uint      `gorm:"primaryKey;autoIncrement" json:"trip_id"`
	DriverID   uint      `json:"driver_id"`
	Driver     Driver    `gorm:"foreignKey:DriverID" json:"driver"`
	CarID      uint      `json:"car_id"`
	Car        Car       `gorm:"foreignKey:CarID" json:"car"`
	CustomerID uint      `json:"customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	StartLat   float64   `gorm:"type:decimal(9,6)" json:"start_lat"`
	StartLon   float64   `gorm:"type:decimal(9,6)" json:"start_lon"`
	EndLat     float64   `gorm:"type:decimal(9,6)" json:"end_lat"`
	EndLon     float64   `gorm:"type:decimal(9,6)" json:"end_lon"`
	StartTime  time.Time `gorm:"type:datetime(8)" json:"start_time"`
	EndTime    time.Time `gorm:"type:datetime(8)" json:"end_time"`
	Cost       float64   `gorm:"type:decimal(10,2)" json:"cost"`
}
