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

	type Query struct {
		WhereClause string `json:"whereClause"`
	}

	//read json request body
	var query Query
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var historicalpo []repository.HistoricalPO
	if query.WhereClause == "" {
		log.Print("whereClause is empty")
		// Fetch customer data from the database
		historicalpo, err = svc.HistoricalpoRepository.GetAll()
	} else {
		log.Print("whereClause is not empty")
		// Fetch customer data from the database based on the whereClause
		historicalpo, err = svc.HistoricalpoRepository.GetWithWhereClause(query.WhereClause)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		// Write the customer data as JSON to the response
		err = json.NewEncoder(w).Encode(historicalpo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
