package repository

/*
CREATE TABLE IF NOT EXISTS pricelists (
    id SERIAL PRIMARY KEY,
    supplier_id VARCHAR(255) NOT NULL,
    sku_id VARCHAR(255) NOT NULL,
    price_per_unit FLOAT NOT NULL,
    stock_available INT NOT NULL
);
*/

// create repository database operation for customer
type Pricelists struct {
	ID             int     `json:"id"`
	SupplierID     string  `json:"supplier_id"`
	SKUID          string  `json:"sku_id"`
	PricePerUnit   float64 `json:"price_per_unit"`
	StockAvailable int     `json:"stock_available"`
}

type PricelistRepository struct {
	Repository *Repository
}

func NewPricelistRepository(repository *Repository) *PricelistRepository {
	return &PricelistRepository{Repository: repository}
}

// Assuming you have a method in your repository like this
func (c *PricelistRepository) GetAll() ([]Pricelists, error) {
	rows, err := c.Repository.DB.Query("SELECT * FROM pricelists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pricelists []Pricelists

	for rows.Next() {
		var cus Pricelists
		err := rows.Scan(&cus.ID, &cus.SupplierID, &cus.SKUID, &cus.PricePerUnit, &cus.StockAvailable)
		if err != nil {
			return nil, err
		}

		pricelists = append(pricelists, cus)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pricelists, nil
}
