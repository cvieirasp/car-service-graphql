package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Order struct {
	db          *sql.DB
	ID          string
	CustomerID  string
	VehicleID   string
	ServiceID   string
	RequestDate string
	Status      string
}

func NewOrder(db *sql.DB) *Order {
	return &Order{
		db: db,
	}
}

func (o *Order) Create(customerID, vehicleID, serviceID, requestDate, status string) (*Order, error) {
	id := uuid.New().String()
	_, err := o.db.Exec("INSERT INTO orders (id, customer_id, vehicle_id, service_id, request_date, status) VALUES ($1, $2, $3, $4, $5, $6)", id, customerID, vehicleID, serviceID, requestDate, status)
	if err != nil {
		return nil, err
	}
	return &Order{ID: id, CustomerID: customerID, VehicleID: vehicleID, ServiceID: serviceID, RequestDate: requestDate, Status: status}, nil
}

func (o *Order) FindAll() ([]Order, error) {
	rows, err := o.db.Query("SELECT id, customer_id, vehicle_id, service_id, request_date, status FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []Order{}
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.VehicleID, &order.ServiceID, &order.RequestDate, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (o *Order) FindByID(id string) (Order, error) {
	var order Order
	if err := o.db.QueryRow("SELECT id, customer_id, vehicle_id, service_id, request_date, status FROM orders WHERE id = $1", id).Scan(&order.ID, &order.CustomerID, &order.VehicleID, &order.ServiceID, &order.RequestDate, &order.Status); err != nil {
		return Order{}, err
	}
	return order, nil
}

func (o *Order) FindByCustomerID(customerID string) ([]Order, error) {
	rows, err := o.db.Query("SELECT id, customer_id, vehicle_id, service_id, request_date, status FROM orders WHERE customer_id = $1", customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []Order{}
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.VehicleID, &order.ServiceID, &order.RequestDate, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (o *Order) FindByVehicleID(vehicleID string) ([]Order, error) {
	rows, err := o.db.Query("SELECT id, customer_id, vehicle_id, service_id, request_date, status FROM orders WHERE vehicle_id = $1", vehicleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []Order{}
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.VehicleID, &order.ServiceID, &order.RequestDate, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (o *Order) FindByServiceID(serviceID string) ([]Order, error) {
	rows, err := o.db.Query("SELECT id, customer_id, vehicle_id, service_id, request_date, status FROM orders WHERE service_id = $1", serviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []Order{}
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.VehicleID, &order.ServiceID, &order.RequestDate, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
