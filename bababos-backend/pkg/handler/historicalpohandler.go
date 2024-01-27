package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prima101112/bababos-backend/repository"
)

type SvcHistoricalpo struct {
	HistoricalpoRepository *repository.HistoricalpoRepository
}

func NewHistoricalpoHandler(con *repository.Repository) *SvcHistoricalpo {
	return &SvcHistoricalpo{
		HistoricalpoRepository: repository.NewHistoricalpoRepository(con),
	}
}

func (svc *SvcHistoricalpo) GetHistoricalposHandler(w http.ResponseWriter, r *http.Request) {

	// Parse query parameters from the request URL
	queryParams := r.URL.Query()
	var err error
	skuid := queryParams.Get("skuid")

	var historicalpo []repository.HistoricalPO
	if skuid == "" {
		log.Print("whereClause is empty")
		// Fetch customer data from the database
		historicalpo, err = svc.HistoricalpoRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Fetch customer data from the database based on the whereClause
		historicalpo, err = svc.HistoricalpoRepository.GetWithWhereClause("sku_id = " + skuid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Write the customer data as JSON to the response
	err = json.NewEncoder(w).Encode(historicalpo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
