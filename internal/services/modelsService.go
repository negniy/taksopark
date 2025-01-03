package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"taksopark/internal/DTO"
	"taksopark/internal/models"

	"gorm.io/gorm"
)

type ModelService struct {
	db *gorm.DB
}

func NewModelService(init_db *gorm.DB) ModelService {
	return ModelService{
		db: init_db,
	}
}

func (s *ModelService) Create(w http.ResponseWriter, r *http.Request) {
	req := new(DTO.CreateModelRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	model := &models.CarModel{
		ModelName:    req.ModelName,
		Manufacturer: req.Manufacturer,
	}

	if err := s.db.Create(model).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusCreated, model)
}

func (s *ModelService) GetAll(w http.ResponseWriter, r *http.Request) {
	var models []models.CarModel
	err := s.db.Find(&models).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusOK, models)
}

func (s *ModelService) Get(w http.ResponseWriter, r *http.Request) {
	var model models.CarModel

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = s.db.Find(&model, id).Error
	if err != nil {
		responseError(w, http.StatusNotFound, err)
		return
	}

	response(w, http.StatusOK, model)
}

func (s *ModelService) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	req := new(DTO.UpdateModelRequest)
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	var model models.CarModel
	err = s.db.Find(&model, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(w, http.StatusNotFound, err)
			return
		}
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	model.ModelName = req.ModelName
	model.Manufacturer = req.Manufacturer

	err = s.db.Save(&model).Error
	if err != nil {
		response(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, nil)
}

func (s *ModelService) UpdateSomething(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var model models.CarModel
	err = s.db.Find(&model, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(w, http.StatusNotFound, err)
			return
		}
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	req := new(DTO.UpdateSomethingModelRequest)
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	switch {
	case req.ModelName != nil:
		model.ModelName = *req.ModelName
	case req.Manufacturer != nil:
		model.Manufacturer = *req.Manufacturer
	}

	err = s.db.Save(&model).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

func (s *ModelService) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = s.db.Delete(models.CarModel{}, id).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusNoContent, nil)
}
