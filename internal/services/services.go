package services

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Service struct {
	Cars      CarService
	Models    ModelService
	Drivers   DriverService
	Customers CustomerService
	Trips     TripService
	Query     QueryService
}

func response(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println(err)
		}
	}
}

func responseError(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{"error :": err.Error()})
}

func NewService(db *gorm.DB) Service {
	return Service{
		Cars:      NewCarService(db),
		Query:     NewQueryService(db),
		Models:    NewModelService(db),
		Drivers:   NewDriverService(db),
		Customers: NewCustomerService(db),
		Trips:     NewTripService(db),
	}
}
