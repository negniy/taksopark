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
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

}

func Run() error {

	service := *services.NewService(db)

	h := http.NewServeMux()
	h.HandleFunc("GET /cars", service.Cars.GetAll)
	h.HandleFunc("GET /cars/{id}", service.Cars.Get)
	h.HandleFunc("PUT /cars/{id}", service.Cars.Update)
	h.HandleFunc("PATCH /cars/{id}", service.Cars.UpdateSomething)
	h.HandleFunc("DELETE /cars/{id}", service.Cars.Delete)

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
	log.Printf("%s %v - %s", "Reseived shutdown signal", sig, "")
	return server.Shutdown(context.Background())
}

func main() {

	InitDB()

	if err := Run(); err != nil {
		log.Fatalf("Server stopped with error: %v", err)
	}

}
