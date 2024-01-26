package seed

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func Seed() {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_DB_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
	)

	fmt.Println("Connection String:", connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	customerSeed(db)
	supplierSeed(db)
	rfqsSeed(db)
	PriceListSeed(db)
	err = insertHistoricalPOData(db, "/app/migrations/seed/historicalpo.csv")
	if err != nil {
		log.Print(err)
	}
}

func customerSeed(db *sql.DB) {
	//clear data first truncate but ignore the errors
	db.Exec(`TRUNCATE TABLE customers`)

	// Seed data
	seedData := []struct {
		CustomerID string
		Address    string
		City       string
		State      string
	}{
		{"M1-ABDI-11", "Taman Sari, Kec. Setu", "Bekasi", "Jawa Barat"},
		{"M1-PUMS-11", "Gunung Sindur", "Bogor", "Jawa Barat"},
		{"M1-STQI-11", "Kel. Rorotan- Cilincing", "Jakarta Utara", "DKI Jakarta"},
		{"M1-SUGP-11", "Gn. Putri Kec. Gn. Putri", "Bekasi", "Jawa Barat"},
		{"M1-SURA-11", "Padurenan, Mustika Jaya", "Bekasi", "Jawa Barat"},
	}

	for _, data := range seedData {
		_, err := db.Exec(`
				INSERT INTO customers (customer_id, address, city, state)
				VALUES ($1, $2, $3, $4)`,
			data.CustomerID, data.Address, data.City, data.State)

		if err != nil {
			log.Print(err)
		}
	}

	fmt.Println("Data customer seeded successfully.")
}

func supplierSeed(db *sql.DB) {
	//clear data first truncate but ignore the errors
	db.Exec(`TRUNCATE TABLE suppliers`)

	type Supplier struct {
		SupplierID string
		Address    string
		City       string
		State      string
	}

	// Seed data
	seedData := []Supplier{
		{"S1-KPS-1", "Kelurahan Klender, Kecamatan Duren Sawit", "Jakarta Timur", "DKI Jakarta"},
		{"S1-HSC-1", "Mangga Dua Sel., Kecamatan Sawah Besar", "Jakarta Pusat", "DKI Jakarta"},
		{"S1-PSB-1", "Tubagus Angke", "Jakarta", "DKI Jakarta"},
		{"S1-SUM-1", "Kedung Waringin, Kec. Tanah Sereal", "Kota Bogor", "Jawa Barat"},
		{"S1-ISB-1", "Penjaringan", "Jakarta Utara", "DKI Jakarta"},
		{"S1-FIX-1", "Godangdia", "Jakarta Pusat", "DKI Jakarta"},
		{"S1-SAM-1", "Bantar Gebang", "Kota Bekasi", "Jawa Barat"},
		{"S1-SSC-1", "Kec. Sawah Besar", "Jakarta Pusat", "DKI Jakarta"},
	}

	for _, data := range seedData {
		_, err := db.Exec(`
			INSERT INTO suppliers (supplier_id, address, city, state)
			VALUES ($1, $2, $3, $4)`,
			data.SupplierID, data.Address, data.City, data.State)

		if err != nil {
			log.Print(err)
		}
	}

	fmt.Println("Data supplier successfully.")
}

func rfqsSeed(db *sql.DB) {
	//clear data first truncate but ignore the errors
	db.Exec(`TRUNCATE TABLE rfqs`)

	type Rfq struct {
		CustomerID    string
		SKUID         string
		Quantity      int
		UnitOfMeasure string
	}
	// Seed data
	seedData := []Rfq{
		{"M1-ABDI-11", "UNP-120", 10, "Batang"},
		{"M1-ABDI-11", "UNP-200", 6, "Batang"},
		{"M1-STQI-11", "PLT-BRDS0230", 100, "Lembar"},
		{"M1-STQI-11", "PLT-SPHC0155", 50, "Lembar"},
		{"M1-STQI-11", "PLT-SPHC0180", 50, "Lembar"},
		{"M1-SURA-11", "PLT-SPHC0120", 170, "Lembar"},
		{"M1-SURA-11", "PLT-SPHC0150", 20, "Lembar"},
		{"M1-SURA-11", "PLT-SPHC0200", 15, "Lembar"},
		{"M1-PUMS-11", "SIK-040040-IBB", 51, "Batang"},
		{"M1-PUMS-11", "SIK-050050-IBB", 72, "Batang"},
		{"M1-PUMS-11", "SIK-060060-IBB", 8, "Batang"},
		{"M1-PUMS-11", "SIK-070070-IBB", 8, "Batang"},
		{"M1-PUMS-11", "SIK-080080-IBB", 10, "Batang"},
		{"M1-PUMS-11", "SIK-120120-IBB", 5, "Batang"},
		{"M1-PUMS-11", "PIP-SCH4080", 2, "Batang"},
		{"M1-SUGP-11", "SIK-100100-IBB", 30, "Batang"},
		{"M1-SUGP-11", "SIK-120120-IBB", 11, "Batang"},
		{"M1-SUGP-11", "SIK-060060-IBB", 50, "Batang"},
		{"M1-SUGP-11", "SIK-070070-IBB", 120, "Batang"},
		{"M1-SUGP-11", "SIK-080080-IBB", 15, "Batang"},
		{"M1-SUGP-11", "WFL-300-GG", 8, "Batang"},
	}

	for _, data := range seedData {
		_, err := db.Exec(`
			INSERT INTO rfqs (customer_id, sku_id, quantity, unit_of_measure)
			VALUES ($1, $2, $3, $4)`,
			data.CustomerID, data.SKUID, data.Quantity, data.UnitOfMeasure)

		if err != nil {
			log.Print(err)
		}
	}

	fmt.Println("Data seeded successfully.")
}

func PriceListSeed(db *sql.DB) {
	//clear data first truncate but ignore the errors
	db.Exec(`TRUNCATE TABLE pricelists`)
	type Pricelist struct {
		SupplierID     string
		SKUID          string
		PricePerUnit   float64
		StockAvailable int
	}

	// Seed data
	seedData := []Pricelist{
		{"S1-FIX-1", "PLT-SPHC1000", 5596216.216, 1},
		{"S1-FIX-1", "PLT-SPHC0400", 1789729.73, 3},
		{"S1-FIX-1", "PLT-SPHC0500", 2390540.541, 5},
		{"S1-FIX-1", "PLT-SPHC0600", 3008783.784, 4},
		{"S1-FIX-1", "PLT-SPHC1200-GG", 3640765.766, 11},
		{"S1-FIX-1", "SIK-100100-IBB", 1066572.973, 30},
		{"S1-FIX-1", "SIK-120120-IBB", 1540090.09, 5},
		{"S1-FIX-1", "SIK-040040-IBB", 166129.7297, 51},
		{"S1-FIX-1", "SIK-050050-IBB", 259720.7207, 72},
		{"S1-FIX-1", "SIK-060060-IBB", 372075.6757, 8},
		{"S1-FIX-1", "SIK-070070-IBB", 506627.027, 8},
		{"S1-HSC-1", "PIP-SCH404", 1559459.459, 3},
		{"S1-HSC-1", "PIP-SCH4060", 2727027.027, 3},
		{"S1-HSC-1", "PIP-SCH4080", 4163063.063, 2},
		{"S1-HSC-1", "PLT-SPHC1200-GG", 3544144.144, 11},
		{"S1-HSC-1", "SIK-100100-IBB", 990990.991, 30},
		{"S1-HSC-1", "SIK-120120-IBB", 1447747.748, 30},
		{"S1-HSC-1", "SIK-040040-IBB", 143694, 51},
		{"S1-HSC-1", "SIK-060060-IBB", 324775, 112},
		{"S1-HSC-1", "SIK-070070-IBB", 443243, 24},
		{"S1-HSC-1", "SIK-080080-IBB", 633783.7838, 24},
		{"S1-HSC-1", "SIK-090090-KS", 804504.5045, 52},
		{"S1-HSC-1", "UNP-200", 2266667, 6},
		{"S1-HSC-1", "PLT-SPHC0155", 518919, 50},
		{"S1-HSC-1", "PLT-SPHC0180", 582702.7027, 50},
		{"S1-HSC-1", "UNP-120", 1968468.468, 10},
		{"S1-ISB-1", "UNP-200", 4387387.387, 6},
		{"S1-ISB-1", "PLT-BRDS0230", 828828.8288, 100},
		{"S1-ISB-1", "PLT-SPHC0120", 400900.9009, 170},
		{"S1-KPS-1", "PLT-SPHC0150", 504504.5045, 20},
		{"S1-KPS-1", "PLT-SPHC0200", 627027.027, 15},
		{"S1-PSB-1", "PLT-SPHC0026", 788288.2883, 10},
		{"S1-PSB-1", "PLT-SPHC0280", 855855.8559, 20},
		{"S1-PSB-1", "SIK-100100-IBB", 1576767.568, 30},
		{"S1-PSB-1", "SIK-120120-IBB", 2366674.775, 12},
		{"S1-SAM-1", "SIK-040040-IBB", 151577.4775, 51},
		{"S1-SAM-1", "SIK-050050-IBB", 475662.1622, 72},
		{"S1-SSC-1", "SIK-060060-IBB", 687067.5676, 112},
		{"S1-SSC-1", "SIK-070070-IBB", 468262.1622, 124},
		{"S1-SSC-1", "SIK-080080-IBB", 1233198.198, 20},
		{"S1-SSC-1", "WFL-300-GG", 7040000, 8},
		{"S1-SSC-1", "PLT-KPL0100", 9864865, 2},
		{"S1-SSC-1", "SIK-050050-IBB", 189189, 72},
		{"S1-SSC-1", "SIK-120120-IBB", 1261261.261, 5},
		{"S1-SSC-1", "SIK-040040-IBB", 144144.1441, 51},
		{"S1-SUM-1", "SIK-050050-IBB", 189189.1892, 72},
	}

	for _, data := range seedData {
		_, err := db.Exec(`
		INSERT INTO pricelists (supplier_id, sku_id, price_per_unit, stock_available)
		VALUES ($1, $2, $3, $4)`,
			data.SupplierID, data.SKUID, data.PricePerUnit, data.StockAvailable)

		if err != nil {
			log.Print(err)
		}
	}

	fmt.Println("Data seeded successfully.")
}

func insertHistoricalPOData(db *sql.DB, filePath string) error {

	type HistoricalPO struct {
		CustomerID       string
		OrderDate        string
		SKUCode          string
		SKUID            string
		SKUName          string
		OrderQuantity    int
		OrderUnit        string
		UnitSellingPrice float64
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Assuming the first row contains the header, use it to get column indices
	header, err := reader.Read()
	if err != nil {
		return err
	}

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		data := make(map[string]string)
		for i, value := range line {
			data[header[i]] = value
		}

		historicalPO := HistoricalPO{
			CustomerID:       data["customer_id"],
			OrderDate:        formatDate(data["order_date"]),
			SKUCode:          data["sku_code"],
			SKUID:            data["sku_id"],
			SKUName:          data["sku_name"],
			OrderQuantity:    atoi(data["order_quantity"]),
			OrderUnit:        data["order_unit"],
			UnitSellingPrice: atof(data["unit_selling_price"]),
		}

		// Insert into the database
		_, err = db.Exec(`
			INSERT INTO historicalpo (customer_id, order_date, sku_code, sku_id, sku_name, order_quantity, order_unit, unit_selling_price)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`, historicalPO.CustomerID, historicalPO.OrderDate, historicalPO.SKUCode, historicalPO.SKUID, historicalPO.SKUName, historicalPO.OrderQuantity, historicalPO.OrderUnit, historicalPO.UnitSellingPrice)

		if err != nil {
			return err
		}
	}

	return nil
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func atof(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func formatDate(s string) string {
	// Parse the input string into a time.Time value
	t, err := time.Parse("02/01/2006", s)
	if err != nil {
		fmt.Println(err)
		log.Fatal("failed seed")
	}

	// Format the time.Time value into the PostgreSQL date format
	output := t.Format("2006-01-02")

	return output
}
