package graph

import "github.com/cvieirasp/car-service-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
	ServiceDB  *database.Service
	CustomerDB *database.Customer
	VehicleDB  *database.Vehicle
	OrderDB    *database.Order
}
