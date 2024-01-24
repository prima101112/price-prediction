package handler

import (
	"encoding/json"
	"net/http"

	"github.com/prima101112/bababos-backend/repository"
)

type SvcCustomer struct {
	*repository.CustomerRepository
}

func NewCustomerhandler(con *repository.Repository) *SvcCustomer {
	return &SvcCustomer{
		CustomerRepository: repository.NewCustomerRepository(con),
	}
}

func (svc *SvcCustomer) GetCustomersHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch customer data from the database based on the customer ID
	customers, err := svc.CustomerRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		// Write the customer data as JSON to the response
		err = json.NewEncoder(w).Encode(customers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
