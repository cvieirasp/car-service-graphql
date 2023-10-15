// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewCategory struct {
	Name string `json:"name"`
}

type NewCustomer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type NewOrder struct {
	CustomerID  string `json:"customerID"`
	VehicleID   string `json:"vehicleID"`
	ServiceID   string `json:"serviceID"`
	RequestDate string `json:"requestDate"`
	Status      string `json:"status"`
}

type NewService struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"categoryID"`
}

type NewVehicle struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}