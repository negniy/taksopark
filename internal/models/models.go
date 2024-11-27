package models

import (
	"time"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
)

type Driver struct {
	DriverID      uint   `gorm:"primaryKey;autoIncrement"`
	FirstName     string `gorm:"size:100"`
	LastName      string `gorm:"size:100"`
	LicenseNumber string `gorm:"uniqueIndex"`
}

type CarModel struct {
	ModelID      uint   `gorm:"primaryKey;autoIncrement"`
	ModelName    string `gorm:"size:100"`
	Manufacturer string `gorm:"size:100"`
}

type Car struct {
	CarID        uint   `gorm:"primaryKey;autoIncrement"`
	LicensePlate string `gorm:"size:100;uniqueIndex"`
	ModelID      uint
	Year         time.Time `gorm:"type:year"`
	Notes        string    `gorm:"type:json"`
}

type Customer struct {
	CustomerID uint   `gorm:"primaryKey;autoIncrement"`
	FirstName  string `gorm:"size:100"`
	LastName   string `gorm:"size:100"`
	Phone      string `gorm:"size:15"`
}

type Trip struct {
	TripID     uint `gorm:"primaryKey;autoIncrement"`
	DriverID   uint
	CarID      uint
	CustomerID uint
	StartLat   float64   `gorm:"type:decimal(9,6)"`
	StartLon   float64   `gorm:"type:decimal(9,6)"`
	EndLat     float64   `gorm:"type:decimal(9,6)"`
	EndLon     float64   `gorm:"type:decimal(9,6)"`
	StartTime  time.Time `gorm:"type:datetime(8)"`
	EndTime    time.Time `gorm:"type:datetime(8)"`
	Cost       float64   `gorm:"type:decimal(10,2)"`
}
