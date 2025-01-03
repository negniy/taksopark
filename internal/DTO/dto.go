package DTO

import (
	"time"

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

type CreateCarRequest struct {
	LicensePlate string `json:"license_plate"`
	ModelID      uint   `json:"model_id"`
	Year         int    `json:"year"`
	Notes        string `json:"notes"`
}

type UpdateCarRequest struct {
	LicensePlate string `json:"license_plate"`
	ModelID      uint   `json:"model_id"`
	Year         int    `json:"year"`
	Notes        string `json:"notes"`
}

type UpdateSomethingCarRequest struct {
	LicensePlate *string `json:"license_plate,omitempty"`
	ModelID      *uint   `json:"model_id,omitempty"`
	Year         *int    `json:"year,omitempty"`
	Notes        *string `json:"notes,omitempty"`
}

type CreateCustomerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type UpdateCustomerRequest struct {
	FirstName string `gorm:"size:100" json:"first_name"`
	LastName  string `gorm:"size:100" json:"last_name"`
	Phone     string `gorm:"size:15" json:"phone"`
}

type UpdateSomethingCustomerRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Phone     *string `json:"phone,omitempty"`
}

type CreateDriverRequest struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	LisenceNumber string `json:"lisence_number"`
}

type UpdateDriverRequest struct {
	FirstName     string `gorm:"size:100" json:"first_name"`
	LastName      string `gorm:"size:100" json:"last_name"`
	LicenseNumber string `gorm:"uniqueIndex" json:"license_number"`
}
type UpdateSomethingDriverRequest struct {
	FirstName     *string `json:"first_name,omitempty"`
	LastName      *string `json:"last_name,omitempty"`
	LisenceNumber *string `json:"lisence_number,omitempty"`
}

type CreateModelRequest struct {
	ModelName    string `json:"model_name"`
	Manufacturer string `json:"manufacturer"`
}
type UpdateModelRequest struct {
	ModelName    string `gorm:"size:100" json:"model_name"`
	Manufacturer string `gorm:"size:100" json:"manufacturer"`
}

type UpdateSomethingModelRequest struct {
	ModelName    *string `json:"model_name,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
}
type CreateTripRequest struct {
	DriverID   uint      `json:"driver_id"`
	CarID      uint      `json:"car_id"`
	CustomerID uint      `json:"customer_id"`
	StartLat   float64   `gorm:"type:decimal(9,6)" json:"start_lat"`
	StartLon   float64   `gorm:"type:decimal(9,6)" json:"start_lon"`
	EndLat     float64   `gorm:"type:decimal(9,6)" json:"end_lat"`
	EndLon     float64   `gorm:"type:decimal(9,6)" json:"end_lon"`
	StartTime  time.Time `gorm:"type:datetime(8)" json:"start_time"`
	EndTime    time.Time `gorm:"type:datetime(8)" json:"end_time"`
	Cost       float64   `gorm:"type:decimal(10,2)" json:"cost"`
}

type UpdateTripRequest struct {
	DriverID   uint      `json:"driver_id"`
	CarID      uint      `json:"car_id"`
	CustomerID uint      `json:"customer_id"`
	StartLat   float64   `json:"start_lat"`
	StartLon   float64   `json:"start_lon"`
	EndLat     float64   `json:"end_lat"`
	EndLon     float64   `json:"end_lon"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Cost       float64   `json:"cost"`
}
type UpdateSomethingTripRequest struct {
	DriverID   *uint      `json:"driver_id,omitempty"`
	CarID      *uint      `json:"car_id,omitempty"`
	CustomerID *uint      `json:"customer_id,omitempty"`
	StartLat   *float64   `json:"start_lat,omitempty"`
	StartLon   *float64   `json:"start_lon,omitempty"`
	EndLat     *float64   `json:"end_lat,omitempty"`
	EndLon     *float64   `json:"end_lon,omitempty"`
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	Cost       *float64   `json:"cost,omitempty"`
}
