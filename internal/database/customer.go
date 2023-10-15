package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Customer struct {
	db      *sql.DB
	ID      string
	Name    string
	Address string
}

func NewCustomer(db *sql.DB) *Customer {
	return &Customer{
		db: db,
	}
}

func (c *Customer) Create(name, address string) (*Customer, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO customers (id, name, address) VALUES ($1, $2, $3)", id, name, address)
	if err != nil {
		return nil, err
	}
	return &Customer{ID: id, Name: name, Address: address}, nil
}

func (c *Customer) FindAll() ([]Customer, error) {
	rows, err := c.db.Query("SELECT id, name, address FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := []Customer{}
	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Address); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *Customer) FindByID(id string) (Customer, error) {
	var customer Customer
	if err := c.db.QueryRow("SELECT id, name, address FROM customers WHERE id = $1", id).Scan(&customer.ID, &customer.Name, &customer.Address); err != nil {
		return Customer{}, err
	}
	return customer, nil
}

func (C *Customer) FindByOrderID(orderID string) (Customer, error) {
	var customer Customer
	if err := C.db.QueryRow("SELECT c.id, c.name, c.address FROM customers c JOIN orders o ON c.id = o.customer_id WHERE o.id = $1", orderID).Scan(&customer.ID, &customer.Name, &customer.Address); err != nil {
		return Customer{}, err
	}
	return customer, nil
}
