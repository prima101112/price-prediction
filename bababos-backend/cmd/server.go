package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prima101112/bababos-backend/migrations/seed"
	"github.com/prima101112/bababos-backend/pkg/handler"
	"github.com/prima101112/bababos-backend/repository"
)

func NewServer() {

	//initialize database repository
	db := repository.NewRepository()
	defer db.DB.Close()

	// initiate customer service
	customerService := handler.NewCustomerhandler(&repository.Repository{DB: db.DB})

	// initiate historicalpo service
	historicalpoService := handler.NewHistoricalpoHandler(&repository.Repository{DB: db.DB})

	// initiate historicalpo service
	pricelistsService := handler.NewPricelistsHandler(&repository.Repository{DB: db.DB})

	r := mux.NewRouter()
	r.HandleFunc("/customers", customerService.GetCustomersHandler).Methods("GET")
	r.HandleFunc("/historicalpo", historicalpoService.GetHistoricalposHandler).Methods("GET")
	r.HandleFunc("/pricelist", pricelistsService.GetPricelistsHandler).Methods("GET")
	http.Handle("/", r)

	// Create a server and specify the address and port to listen on
	serverAddr := ":8080" // You can change the port as needed
	server := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	fmt.Printf("Server listening on %s...\n", serverAddr)

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func main() {
	//enable in dev mode always truncate before seeding
	// seeding not needed in staging or production
	seed.Seed()
	NewServer()
}
