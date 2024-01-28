package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prima101112/bababos-backend/repository"
)

type SvcSuggestion struct {
	SuggestionRepository *repository.SuggestionRepository
}

func NewSuggestionHandler(con *repository.Repository) *SvcSuggestion {
	return &SvcSuggestion{
		SuggestionRepository: repository.NewSuggestionRepository(con),
	}
}

func (svc *SvcSuggestion) GetSuggestionHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters from the request URL
	queryParams := r.URL.Query()
	var err error
	skuid := queryParams.Get("sku_id")
	quantity := queryParams.Get("quantity")

	var suggestion repository.Suggestion

	// Fetch customer data from the database based on the whereClause
	suggestion, err = svc.SuggestionRepository.PredictSalesBySku(skuid, quantity)
	if err != nil {
		//TODO should handle properly later
		log.Println(err)
		err = json.NewEncoder(w).Encode(suggestion)
		return
	}

	// Write the customer data as JSON to the response
	err = json.NewEncoder(w).Encode(suggestion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
