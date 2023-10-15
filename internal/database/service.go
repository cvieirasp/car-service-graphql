package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Service struct {
	db          *sql.DB
	ID          string
	Description string
	Price       float64
	CategoryID  string
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Create(description string, price float64, categoryID string) (*Service, error) {
	id := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO services (id, description, price, category_id) VALUES ($1, $2, $3, $4)", id, description, price, categoryID)
	if err != nil {
		return nil, err
	}
	return &Service{ID: id, Description: description, Price: price, CategoryID: categoryID}, nil
}

func (s *Service) FindAll() ([]Service, error) {
	rows, err := s.db.Query("SELECT id, description, price, category_id FROM services")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	services := []Service{}
	for rows.Next() {
		var service Service
		if err := rows.Scan(&service.ID, &service.Description, &service.Price, &service.CategoryID); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (s *Service) FindByID(id string) (Service, error) {
	var service Service
	if err := s.db.QueryRow("SELECT id, description, price, category_id FROM services WHERE id = $1", id).Scan(&service.ID, &service.Description, &service.Price, &service.CategoryID); err != nil {
		return Service{}, err
	}
	return service, nil
}

func (s *Service) FindByCategoryID(categoryID string) ([]Service, error) {
	rows, err := s.db.Query("SELECT id, description, price, category_id FROM services WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	services := []Service{}
	for rows.Next() {
		var service Service
		if err := rows.Scan(&service.ID, &service.Description, &service.Price, &service.CategoryID); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (s *Service) FindByOrderID(orderID string) (Service, error) {
	var service Service
	if err := s.db.QueryRow("SELECT s.id, s.description, s.price, s.category_id FROM services s JOIN orders o ON s.id = o.service_id WHERE o.id = $1", orderID).Scan(&service.ID, &service.Description, &service.Price, &service.CategoryID); err != nil {
		return Service{}, err
	}
	return service, nil
}
