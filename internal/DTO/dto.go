package dto

import (
	"taksopark/internal/models"
	"time"
)

type CarWithModel struct {
	CarID        uint
	LicensePlate string
	Model        models.CarModel
	Year         time.Time
	Notes        string
}
