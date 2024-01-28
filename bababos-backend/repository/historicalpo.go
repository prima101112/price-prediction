package repository

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
	ID               int     `json:"id"`
	CustomerID       string  `json:"customer_id"`
	OrderDate        string  `json:"order_date"`
	SKUCode          string  `json:"sku_code"`
	SKUID            string  `json:"sku_id"`
	SKUName          string  `json:"sku_name"`
	OrderQuantity    int     `json:"order_quantity"`
	OrderUnit        string  `json:"order_unit"`
	UnitSellingPrice float64 `json:"unit_selling_price"`
}

type ResponsePrediction struct {
	SKUID           string
	SKUName         string
	SalesPO         SalesPO
	PricePrediction float64
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
