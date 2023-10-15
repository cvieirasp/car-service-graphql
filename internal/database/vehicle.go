package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Vehicle struct {
	db    *sql.DB
	ID    string
	Brand string
	Model string
	Year  int
}

func NewVehicle(db *sql.DB) *Vehicle {
	return &Vehicle{
		db: db,
	}
}

func (v *Vehicle) Create(brand, model string, year int) (*Vehicle, error) {
	id := uuid.New().String()
	_, err := v.db.Exec("INSERT INTO vehicles (id, brand, model, year) VALUES ($1, $2, $3, $4)", id, brand, model, year)
	if err != nil {
		return nil, err
	}
	return &Vehicle{ID: id, Brand: brand, Model: model, Year: year}, nil
}

func (v *Vehicle) FindAll() ([]Vehicle, error) {
	rows, err := v.db.Query("SELECT id, brand, model, year FROM vehicles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vehicles := []Vehicle{}
	for rows.Next() {
		var vehicle Vehicle
		if err := rows.Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.Year); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

func (v *Vehicle) FindByID(id string) (Vehicle, error) {
	var vehicle Vehicle
	if err := v.db.QueryRow("SELECT id, brand, model, year FROM vehicles WHERE id = $1", id).Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.Year); err != nil {
		return Vehicle{}, err
	}
	return vehicle, nil
}

func (C *Vehicle) FindByOrderID(orderID string) (Vehicle, error) {
	var vehicle Vehicle
	if err := C.db.QueryRow("SELECT v.id, v.brand, v.model, v.year FROM vehicles v JOIN orders o ON v.id = o.vehicle_id WHERE o.id = $1", orderID).Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.Year); err != nil {
		return Vehicle{}, err
	}
	return vehicle, nil
}
