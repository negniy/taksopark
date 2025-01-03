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

type CarService struct {
	db *gorm.DB
}

func NewCarService(init_db *gorm.DB) CarService {
	return CarService{
		db: init_db,
	}
}

func (s *CarService) Create(w http.ResponseWriter, r *http.Request) {
	req := new(DTO.CreateCarRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	car := &models.Car{
		LicensePlate: req.LicensePlate,
		ModelID:      req.ModelID,
		Year:         uint(req.Year),
		Notes:        req.Notes,
	}

	if err := s.db.Create(car).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusCreated, car)
}

func (s *CarService) Get(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var car models.Car
	err = s.db.Preload("Model").First(&car, id).Error

	switch {
	case err == nil:
		response(w, http.StatusOK, car)
	case errors.Is(err, gorm.ErrRecordNotFound):
		responseError(w, http.StatusNotFound, errors.New("record not found"))
	default:
		responseError(w, http.StatusInternalServerError, err)
	}
}

func (s *CarService) GetAll(w http.ResponseWriter, r *http.Request) {
	var cars []models.Car
	err := s.db.Preload("Model").Find(&cars).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, cars)
}

func (s *CarService) Update(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	req := new(DTO.UpdateCarRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	var car models.Car
	if err := s.db.Preload("Model").First(&car, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(w, http.StatusNotFound, errors.New("car not found"))
			return
		}
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	car.LicensePlate = req.LicensePlate
	car.ModelID = req.ModelID
	car.Notes = req.Notes
	car.Year = uint(req.Year)

	if err := s.db.Save(&car).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, car)
}

func (s *CarService) UpdateSomething(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	req := new(DTO.UpdateSomethingCarRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	var car models.Car
	err = s.db.First(&car, id).Error
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	switch {
	case req.LicensePlate != nil:
		car.LicensePlate = *req.LicensePlate
	case req.ModelID != nil:
		car.ModelID = *req.ModelID
	case req.Year != nil:
		car.Year = uint(*req.Year)
	case req.Notes != nil:
		car.Notes = *req.Notes
	}

	err = s.db.Preload("Model").Save(&car).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

func (s *CarService) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	if err = s.db.Delete(models.Car{}, id).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusNoContent, nil)
}
