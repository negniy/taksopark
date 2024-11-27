package main

import (
	"fmt"
	"log"
	"taksopark/internal/dto"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func query1(year time.Time) {
	var cars []dto.CarWithModel

	query := `
	select car_id, license_plate, year, model_name, manufacturer
	from
	cars inner join car_models cm on cars.model_id = cm.model_id
	where year = ?
	`
	err := db.Raw(query, year).Scan(&cars).Error
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
		return
	}

	for _, car := range cars {
		fmt.Print(car)
	}
}

func InitDB() {
	dbconnect := "root:1234@tcp(127.0.0.1:3306)/taksopark?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dbconnect), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

}

func main() {

	InitDB()

}
