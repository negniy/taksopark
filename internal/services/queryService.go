package services

import (
	"errors"
	"net/http"
	"strconv"
	"taksopark/internal/DTO"

	"gorm.io/gorm"
)

type QueryService struct {
	db *gorm.DB
}

func NewQueryService(init_db *gorm.DB) *QueryService {
	return &QueryService{
		db: init_db,
	}
}

func (q *QueryService) CarOfYear(w http.ResponseWriter, r *http.Request) {

	var res []DTO.CarWithModel
	yearString := r.PathValue("year")
	year, err := strconv.Atoi(yearString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid year"))
		return
	}

	query := q.db.Raw(`
	select car_id, license_plate, cars.year, model_name, manufacturer
	from
	cars inner join car_models cm on cars.model_id = cm.model_id
	where cars.year = ?
	`, year).Scan(&res)

	if query.Error != nil {
		responseError(w, http.StatusBadRequest, query.Error)
		return
	}

	response(w, http.StatusOK, res)
}

func (q *QueryService) DriverTripCounter(w http.ResponseWriter, r *http.Request) {

	var res []DTO.PersonCount

	query := q.db.Raw(`
	select first_name, last_name, count(trips.trip_id) count
	from drivers
	left join trips on trips.driver_id=drivers.driver_id
	group by drivers.driver_id
	order by count desc
	`).Scan(&res)

	if query.Error != nil {
		responseError(w, http.StatusBadRequest, query.Error)
		return
	}

	response(w, http.StatusOK, res)
}

func (q *QueryService) DriverTripAutoCounter(w http.ResponseWriter, r *http.Request) {

	var res []DTO.DriverCount

	query := q.db.Raw(`
	select d.first_name, d.last_name, c.license_plate, count(t.trip_id) count
	from
	drivers d left join trips t on t.driver_id=d.driver_id
	inner join cars c on c.car_id = t.car_id
	group by c.car_id, d.first_name, d.last_name, c.license_plate
	`).Scan(&res)

	if query.Error != nil {
		responseError(w, http.StatusBadRequest, query.Error)
		return
	}

	response(w, http.StatusOK, res)
}

func (q *QueryService) ClientTripMoreThan(w http.ResponseWriter, r *http.Request) {

	var res []DTO.PersonCount
	nString := r.PathValue("n")
	n, err := strconv.Atoi(nString)
	if err != nil {
		responseError(w, http.StatusBadRequest, errors.New("invalid n"))
		return
	}

	query := q.db.Raw(`
	select c.first_name, c.last_name, count(t.trip_id) as count
	from customers c
	left join trips t on c.customer_id = t.customer_id
	group by c.customer_id
	having count(t.trip_id) > ?
	`, n).Scan(&res)

	if query.Error != nil {
		responseError(w, http.StatusBadRequest, query.Error)
		return
	}

	response(w, http.StatusOK, res)
}

func (q *QueryService) BestDrivers(w http.ResponseWriter, r *http.Request) {

	var res []DTO.Person

	query := q.db.Raw(`
	select d.first_name, d.last_name
	from drivers d
	join (
		select driver_id, count(trip_id) as trip_count
		from trips
		group by driver_id
	) t on d.driver_id = t.driver_id
	where t.trip_count = (
		select max(trip_count)
		from (
			select count(trip_id) as trip_count
			from trips
			group by driver_id
		) as max_trips
	)
	`).Scan(&res)

	if query.Error != nil {
		responseError(w, http.StatusBadRequest, query.Error)
		return
	}

	response(w, http.StatusOK, res)
}

func (q *QueryService) Statistic(w http.ResponseWriter, r *http.Request) {

	res := new(DTO.Statistic)

	query := q.db.Raw(`
	select 
		min(timestampdiff(minute, start_time, end_time)) as min,
		avg(timestampdiff(minute, start_time, end_time)) as avg,
		max(timestampdiff(minute, start_time, end_time)) as max
	from trips;
	`).Scan(&res)

	if query.Error != nil {
		responseError(w, http.StatusBadRequest, query.Error)
		return
	}

	response(w, http.StatusOK, res)
}
