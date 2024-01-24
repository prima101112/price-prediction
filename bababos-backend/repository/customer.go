package repository

// create repository database operation for customer
type Customer struct {
	ID         int    `json:"id"`
	CustomerID string `json:"customer_id"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
}

type CustomerRepository struct {
	Repository *Repository
}

func NewCustomerRepository(repository *Repository) *CustomerRepository {
	return &CustomerRepository{Repository: repository}
}

func (c *CustomerRepository) GetById(id int) (*Customer, error) {
	cus := &Customer{}
	c.Repository.DB.QueryRow("SELECT * FROM customers WHERE customer_id = ?", id).Scan(&cus.CustomerID, &cus.Address, &cus.City, &cus.State)
	return cus, nil
}

// Assuming you have a method in your repository like this
func (c *CustomerRepository) GetAll() ([]Customer, error) {
	rows, err := c.Repository.DB.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var cus Customer
		err := rows.Scan(&cus.ID, &cus.CustomerID, &cus.Address, &cus.City, &cus.State)
		if err != nil {
			return nil, err
		}

		customers = append(customers, cus)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
