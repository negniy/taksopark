package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"taksopark/internal/models"

	"gorm.io/gorm"
)

type DriverService struct {
	db *gorm.DB
}

func NewDriverService(init_db *gorm.DB) DriverService {
	return DriverService{
		db: init_db,
	}
}

func (s *DriverService) GetAll(w http.ResponseWriter, r *http.Request) {
	var drivers []models.Driver
	err := s.db.Find(&drivers).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusOK, drivers)
}

func (s *DriverService) Get(w http.ResponseWriter, r *http.Request) {
	var driver models.Driver

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = s.db.Find(&driver, id).Error
	if err != nil {
		responseError(w, http.StatusNotFound, err)
		return
	}

	response(w, http.StatusOK, driver)
}

type UpdateDriverRequest struct {
	FirstName     string `gorm:"size:100" json:"first_name"`
	LastName      string `gorm:"size:100" json:"last_name"`
	LicenseNumber string `gorm:"uniqueIndex" json:"license_number"`
}

func (s *DriverService) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	req := new(UpdateDriverRequest)
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	var driver models.Driver
	err = s.db.Find(&driver, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(w, http.StatusNotFound, err)
			return
		}
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	driver.FirstName = req.FirstName
	driver.LastName = req.LastName
	driver.LicenseNumber = req.LicenseNumber

	err = s.db.Save(&driver).Error
	if err != nil {
		response(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, nil)
}

type UpdateSomethingDriverRequest struct {
	FirstName     *string `json:"first_name,omitempty"`
	LastName      *string `json:"last_name,omitempty"`
	LicenseNumber *string `json:"phone,omitempty"`
}

func (s *DriverService) UpdateSomething(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var driver models.Driver
	err = s.db.Find(&driver, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(w, http.StatusNotFound, err)
			return
		}
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	req := new(UpdateSomethingDriverRequest)
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	switch {
	case req.FirstName != nil:
		driver.FirstName = *req.FirstName
	case req.LastName != nil:
		driver.LastName = *req.LastName
	case req.LicenseNumber != nil:
		driver.LicenseNumber = *req.LicenseNumber
	}

	err = s.db.Save(&driver).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

func (s *DriverService) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = s.db.Delete(models.Driver{}, id).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusNoContent, nil)
}
