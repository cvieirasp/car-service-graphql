package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cvieirasp/car-service-graphql/graph"
	"github.com/cvieirasp/car-service-graphql/internal/database"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/car_service?sslmode=disable")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	serviceDB := database.NewService(db)
	customerDB := database.NewCustomer(db)
	vehicleDB := database.NewVehicle(db)
	orderDB := database.NewOrder(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
		ServiceDB:  serviceDB,
		CustomerDB: customerDB,
		VehicleDB:  vehicleDB,
		OrderDB:    orderDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
