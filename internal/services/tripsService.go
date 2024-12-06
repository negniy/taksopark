package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"taksopark/internal/models"
	"time"

	"strconv"

	"gorm.io/gorm"
)

type TripService struct {
	db *gorm.DB
}

func NewTripService(init_db *gorm.DB) TripService {
	return TripService{
		db: init_db,
	}
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

func (s *TripService) Create(w http.ResponseWriter, r *http.Request) {
	req := new(CreateTripRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	trip := &models.Trip{
		DriverID:   req.DriverID,
		CarID:      req.CarID,
		CustomerID: req.CustomerID,
		StartLat:   req.StartLat,
		StartLon:   req.StartLon,
		EndLat:     req.EndLat,
		EndLon:     req.EndLon,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Cost:       req.Cost,
	}

	if err := s.db.Create(trip).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusCreated, trip)
}

func (s *TripService) Get(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var trip models.Trip
	err = s.db.Preload("Customer").Preload("Driver").Preload("Car").Preload("Car.Model").First(&trip, id).Error

	switch {
	case err == nil:
		response(w, http.StatusOK, trip)
	case errors.Is(err, gorm.ErrRecordNotFound):
		responseError(w, http.StatusNotFound, errors.New("record not found"))
	default:
		responseError(w, http.StatusInternalServerError, err)
	}
}

func (s *TripService) GetAll(w http.ResponseWriter, r *http.Request) {
	var trips []models.Trip
	err := s.db.Preload("Customer").Preload("Driver").Preload("Car").Preload("Car.Model").Find(&trips).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, trips)
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

func (s *TripService) Update(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	req := new(UpdateTripRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	var trip models.Trip
	if err := s.db.Preload("Customer").Preload("Driver").Preload("Car").First(&trip, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(w, http.StatusNotFound, errors.New("trip not found"))
			return
		}
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	trip.DriverID = req.DriverID
	trip.CarID = req.CarID
	trip.CustomerID = req.CustomerID
	trip.StartLat = req.StartLat
	trip.StartLon = req.StartLon
	trip.EndLat = req.EndLat
	trip.EndLon = req.EndLon
	trip.StartTime = req.StartTime
	trip.EndTime = req.EndTime
	trip.Cost = req.Cost

	if err := s.db.Save(&trip).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, trip)
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

func (s *TripService) UpdateSomething(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	req := new(UpdateSomethingTripRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	var trip models.Trip
	err = s.db.First(&trip, id).Error
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	switch {
	case req.DriverID != nil:
		trip.DriverID = *req.DriverID
	case req.CarID != nil:
		trip.CarID = *req.CarID
	case req.CustomerID != nil:
		trip.CustomerID = *req.CustomerID
	case req.StartLon != nil:
		trip.StartLon = *req.StartLon
	case req.StartLat != nil:
		trip.StartLat = *req.StartLat
	case req.EndLat != nil:
		trip.EndLat = *req.EndLat
	case req.EndLon != nil:
		trip.EndLon = *req.EndLon
	case req.StartTime != nil:
		trip.StartTime = *req.StartTime
	case req.EndTime != nil:
		trip.EndTime = *req.EndTime
	case req.Cost != nil:
		trip.Cost = *req.Cost
	}

	err = s.db.Preload("Customer").Preload("Driver").Preload("Car").Save(&trip).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

func (s *TripService) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	if err = s.db.Delete(models.Trip{}, id).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusNoContent, nil)
}
