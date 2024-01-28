package repository

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/prima101112/bababos-backend/pkg/algorithm"
)

// CREATE TABLE IF NOT EXISTS historypo (
//     id SERIAL PRIMARY KEY,
//     customer_id VARCHAR(255) NOT NULL,
//     order_date DATE NOT NULL,
//     sku_id VARCHAR(20) NOT NULL,
//     order_quantity INT NOT NULL,
//     order_unit VARCHAR(12) NOT NULL,
//     unit_selling_price FLOAT NOT NULL
// );

// create repository database operation for customer
type Suggestion struct {
	Message               string         `json:"message"`
	SKUID                 string         `json:"sku_id"`
	SKUName               string         `json:"sku_name"`
	LowestPrice           float64        `json:"lowest_price"`
	HigestPrice           float64        `json:"higest_price"`
	AveragePrice          float64        `json:"average_price"`
	MedianPrice           float64        `json:"median_price"`
	LinearRegressionPrice float64        `json:"linear_regression_price"`
	Historicel_po_data    []HistoricalPO `json:"historical_po_data"`
}

type SalesPO struct {
	OrderQuantity []float64
	Price         []float64
}

type SuggestionRepository struct {
	Repository *Repository
}

func NewSuggestionRepository(repository *Repository) *SuggestionRepository {
	return &SuggestionRepository{Repository: repository}
}

// Assuming you have a method in your repository like this
func (c *SuggestionRepository) GetWithWhereClause(whereClause string) ([]HistoricalPO, error) {
	rows, err := c.Repository.DB.Query("SELECT * FROM historicalpo WHERE " + whereClause)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historicalpos []HistoricalPO

	for rows.Next() {
		var cus HistoricalPO
		err := rows.Scan(&cus.ID, &cus.CustomerID, &cus.OrderDate, &cus.SKUCode, &cus.SKUID, &cus.SKUName, &cus.OrderQuantity, &cus.OrderUnit, &cus.UnitSellingPrice)
		if err != nil {
			return nil, err
		}

		historicalpos = append(historicalpos, cus)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return historicalpos, nil
}

func (c *SuggestionRepository) PredictSalesBySku(sku string, quantity string) (Suggestion, error) {
	var sugestion Suggestion
	sugestion.Message = "error predict sales by sku"
	//validate sku format no space, no special character
	sku = removeSpace(sku)
	quantity = removeSpace(quantity)
	qty, _ := strconv.Atoi(quantity)
	log.Print(sku)

	hposs, _ := c.GetAllHistoricalpo() //ignore error for now
	sugestion.Historicel_po_data = hposs

	if sku == "" {
		sugestion.Message = "sku cannot be empty"
		err := errors.New("sku cannot be empty")
		return sugestion, err
	}

	//get data from price from table pricelist
	var pricelists []Pricelists
	var pricelist Pricelists
	var prices []float64
	rowsprice, err := c.Repository.DB.Query("SELECT price_per_unit FROM pricelists WHERE sku_id = '" + sku + "'")
	if err != nil {
		sugestion.Message = "error getting data from pricelist"
		return sugestion, err
	}
	defer rowsprice.Close()
	for rowsprice.Next() {
		err := rowsprice.Scan(&pricelist.PricePerUnit)
		if err != nil {
			return sugestion, err
		}
		pricelists = append(pricelists, pricelist)
		prices = append(prices, pricelist.PricePerUnit)
	}

	if qty == 0 {
		qty = 10
	}
	// TODO should avoid select * from table in the future
	rows, err := c.Repository.DB.Query("SELECT * FROM historicalpo WHERE sku_id = '" + sku + "'")
	if err != nil {
		sugestion.Message = "error getting data from historicalpo"
		return sugestion, err
	}
	defer rows.Close()
	i := 0
	var historicalpo HistoricalPO
	var historicalpos []HistoricalPO
	var salespo SalesPO
	for rows.Next() {
		err := rows.Scan(&historicalpo.ID, &historicalpo.CustomerID, &historicalpo.OrderDate, &historicalpo.SKUCode, &historicalpo.SKUID, &historicalpo.SKUName, &historicalpo.OrderQuantity, &historicalpo.OrderUnit, &historicalpo.UnitSellingPrice)
		if err != nil {
			return sugestion, err
		}
		salespo.OrderQuantity = append(salespo.OrderQuantity, float64(historicalpo.OrderQuantity))
		salespo.Price = append(salespo.Price, historicalpo.UnitSellingPrice)
		historicalpos = append(historicalpos, historicalpo)
		i++
	}

	//create linear regression
	r, err := algorithm.NewRegression(salespo.Price, salespo.OrderQuantity)
	if err != nil {
		sugestion.Message = "error creating linear regression"
		return sugestion, err
	}

	//predict sales using linear regression
	pred, err := algorithm.PredictedSales(r, qty)
	if err != nil {
		sugestion.Message = "error predicting sales using linear regression"
		return sugestion, err
	}

	sugestion.Message = "success"
	sugestion.SKUID = sku
	sugestion.SKUName = historicalpo.SKUName
	sugestion.AveragePrice = algorithm.Average(prices)
	sugestion.MedianPrice = algorithm.Median(prices)
	sugestion.LowestPrice = algorithm.Lowest(prices)
	sugestion.HigestPrice = algorithm.Highest(prices)
	sugestion.LinearRegressionPrice = pred
	sugestion.Historicel_po_data = historicalpos

	return sugestion, nil
}

// Assuming you have a method in your repository like this
func (c *SuggestionRepository) GetAllHistoricalpo() ([]HistoricalPO, error) {
	rows, err := c.Repository.DB.Query("SELECT * FROM historicalpo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historicalpos []HistoricalPO

	for rows.Next() {
		var cus HistoricalPO
		err := rows.Scan(&cus.ID, &cus.CustomerID, &cus.OrderDate, &cus.SKUCode, &cus.SKUID, &cus.SKUName, &cus.OrderQuantity, &cus.OrderUnit, &cus.UnitSellingPrice)
		if err != nil {
			return nil, err
		}

		historicalpos = append(historicalpos, cus)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return historicalpos, nil
}

// remove space function
func removeSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}
