package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"taksopark/internal/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dbconnect := "root:1234@tcp(127.0.0.1:3306)/taksopark?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dbconnect), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection error to database: %v", err)
	}

}

func Run() error {

	service := services.NewService(db)

	h := http.NewServeMux()
	h.HandleFunc("POST /cars", service.Cars.Create)
	h.HandleFunc("GET /cars", service.Cars.GetAll)
	h.HandleFunc("GET /cars/{id}", service.Cars.Get)
	h.HandleFunc("PUT /cars/{id}", service.Cars.Update)
	h.HandleFunc("PATCH /cars/{id}", service.Cars.UpdateSomething)
	h.HandleFunc("DELETE /cars/{id}", service.Cars.Delete)

	h.HandleFunc("POST /customers", service.Customers.Create)
	h.HandleFunc("GET /customers", service.Customers.GetAll)
	h.HandleFunc("GET /customers/{id}", service.Customers.Get)
	h.HandleFunc("PUT /customers/{id}", service.Customers.Update)
	h.HandleFunc("PATCH /customers/{id}", service.Customers.UpdateSomething)
	h.HandleFunc("DELETE /customers/{id}", service.Customers.Delete)

	h.HandleFunc("POST /models", service.Models.Create)
	h.HandleFunc("GET /models", service.Models.GetAll)
	h.HandleFunc("GET /models/{id}", service.Models.Get)
	h.HandleFunc("PUT /models/{id}", service.Models.Update)
	h.HandleFunc("PATCH /models/{id}", service.Models.UpdateSomething)
	h.HandleFunc("DELETE /models/{id}", service.Models.Delete)

	h.HandleFunc("POST /drivers", service.Drivers.Create)
	h.HandleFunc("GET /drivers", service.Drivers.GetAll)
	h.HandleFunc("GET /drivers/{id}", service.Drivers.Get)
	h.HandleFunc("PUT /drivers/{id}", service.Drivers.Update)
	h.HandleFunc("PATCH /drivers/{id}", service.Drivers.UpdateSomething)
	h.HandleFunc("DELETE /drivers/{id}", service.Drivers.Delete)

	h.HandleFunc("POST /trips", service.Trips.Create)
	h.HandleFunc("GET /trips", service.Trips.GetAll)
	h.HandleFunc("GET /trips/{id}", service.Trips.Get)
	h.HandleFunc("PUT /trips/{id}", service.Trips.Update)
	h.HandleFunc("PATCH /trips/{id}", service.Trips.UpdateSomething)
	h.HandleFunc("DELETE /trips/{id}", service.Trips.Delete)

	h.HandleFunc("GET /cars/year/{year}", service.Query.CarOfYear)
	h.HandleFunc("GET /drivers/count", service.Query.DriverTripCounter)
	h.HandleFunc("GET /drivers/autocount", service.Query.DriverTripAutoCounter)
	h.HandleFunc("GET /clients/trips/{n}", service.Query.ClientTripMoreThan)
	h.HandleFunc("GET /drivers/best", service.Query.BestDrivers)
	h.HandleFunc("GET /statistics", service.Query.Statistic)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: h,
	}

	go func() {
		log.Printf("run server: http://localhostlocalhost:8080")
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("error when listen and serve: %s", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	defer signal.Stop(ch)
	sig := <-ch
	log.Printf("%s - %v", "Reseived shutdown signal", sig)
	return server.Shutdown(context.Background())
}

func main() {

	InitDB()

	if err := Run(); err != nil {
		log.Fatalf("Server stopped with error: %v", err)
	}

}
