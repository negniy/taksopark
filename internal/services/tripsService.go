package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"taksopark/internal/DTO"
	"taksopark/internal/models"

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

func (s *TripService) Create(w http.ResponseWriter, r *http.Request) {
	req := new(DTO.CreateTripRequest)
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

func (s *TripService) Update(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	req := new(DTO.UpdateTripRequest)
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

func (s *TripService) UpdateSomething(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	req := new(DTO.UpdateSomethingTripRequest)
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
