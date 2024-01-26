package handler

import (
	"encoding/json"
	"net/http"

	"github.com/prima101112/bababos-backend/repository"
)

type SvcPricelist struct {
	PricelistRepository *repository.HistoricalpoRepository
}

func NewPricelistsHandler(con *repository.Repository) *SvcHistoricalpo {
	return &SvcHistoricalpo{
		HistoricalpoRepository: repository.NewHistoricalpoRepository(con),
	}
}

func (svc *SvcHistoricalpo) GetPricelistsHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch customer data from the database based on the customer ID
	customers, err := svc.HistoricalpoRepository.GetAll()
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
