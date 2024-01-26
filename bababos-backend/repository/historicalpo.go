package repository

import (
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
type HistoricalPO struct {
	ID               int
	CustomerID       string
	OrderDate        string
	SKUCode          string
	SKUID            string
	SKUName          string
	OrderQuantity    int
	OrderUnit        string
	UnitSellingPrice float64
}

type ResponsePrediction struct {
	SKUID           string
	SKUName         string
	SalesPO         SalesPO
	PricePrediction float64
}

type SalesPO struct {
	Time  []float64
	Price []float64
}

type HistoricalpoRepository struct {
	Repository *Repository
}

func NewHistoricalpoRepository(repository *Repository) *HistoricalpoRepository {
	return &HistoricalpoRepository{Repository: repository}
}

// Assuming you have a method in your repository like this
func (c *HistoricalpoRepository) GetWithWhereClause(whereClause string) ([]HistoricalPO, error) {
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

// Assuming you have a method in your repository like this
func (c *HistoricalpoRepository) GetAll() ([]HistoricalPO, error) {
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

func (c *HistoricalpoRepository) PredictSalesBySku(sku string) (ResponsePrediction, error) {

	var responsePrediction ResponsePrediction
	//validate sku format no space, no special character
	sku = removeSpace(sku)

	rows, err := c.Repository.DB.Query("SELECT sku_name, order_quantity, unit_selling_price FROM historicalpo WHERE sku_id = " + sku + " ORDER BY order_date DESC")
	if err != nil {
		return responsePrediction, err
	}
	defer rows.Close()
	i := 0
	var historicalpo HistoricalPO
	var salespo SalesPO
	for rows.Next() {
		err := rows.Scan(&historicalpo.SKUName, &historicalpo.OrderQuantity, &historicalpo.UnitSellingPrice)
		if err != nil {
			return responsePrediction, err
		}
		salespo.Time[i] = float64(historicalpo.OrderQuantity)
		salespo.Price[i] = historicalpo.UnitSellingPrice

		i++
	}

	r, err := algorithm.NewRegression(salespo.Time, salespo.Price)
	if err != nil {
		return responsePrediction, err
	}
	pred, err := algorithm.PredictedSales(r, len(salespo.Time)+1)
	if err != nil {
		return responsePrediction, err
	}

	responsePrediction.SKUID = sku
	responsePrediction.SKUName = historicalpo.SKUName
	responsePrediction.SalesPO = salespo
	responsePrediction.PricePrediction = pred

	return responsePrediction, nil
}

// remove space function
func removeSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}
