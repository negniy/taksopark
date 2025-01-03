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

type CustomerService struct {
	db *gorm.DB
}

func NewCustomerService(init_db *gorm.DB) CustomerService {
	return CustomerService{
		db: init_db,
	}
}

func (s *CustomerService) Create(w http.ResponseWriter, r *http.Request) {
	req := new(DTO.CreateCustomerRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	customer := &models.Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
	}

	if err := s.db.Create(customer).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusCreated, customer)
}

func (s *CustomerService) GetAll(w http.ResponseWriter, r *http.Request) {
	var customers []models.Customer
	err := s.db.Find(&customers).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusOK, customers)
}

func (s *CustomerService) Get(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer

	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = s.db.Find(&customer, id).Error
	if err != nil {
		responseError(w, http.StatusNotFound, err)
		return
	}

	response(w, http.StatusOK, customer)
}

func (s *CustomerService) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	req := new(DTO.UpdateCustomerRequest)
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	var customer models.Customer
	err = s.db.Find(&customer, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(w, http.StatusNotFound, err)
			return
		}
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	customer.FirstName = req.FirstName
	customer.LastName = req.LastName
	customer.Phone = req.Phone

	err = s.db.Save(&customer).Error
	if err != nil {
		response(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, nil)
}

func (s *CustomerService) UpdateSomething(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var customer models.Customer
	err = s.db.Find(&customer, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(w, http.StatusNotFound, err)
			return
		}
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	req := new(DTO.UpdateSomethingCustomerRequest)
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	switch {
	case req.FirstName != nil:
		customer.FirstName = *req.FirstName
	case req.LastName != nil:
		customer.LastName = *req.LastName
	case req.Phone != nil:
		customer.Phone = *req.Phone
	}

	err = s.db.Save(&customer).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

func (s *CustomerService) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	err = s.db.Delete(models.Customer{}, id).Error
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusNoContent, nil)
}
