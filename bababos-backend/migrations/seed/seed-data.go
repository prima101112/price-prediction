package seed

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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
		log.Fatal(err)
	}
	defer db.Close()

	customerSeed(db)
	supplierSeed(db)
	rfqsSeed(db)
	PriceListSeed(db)
	historicalPOSeed(db)

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
			log.Fatal(err)
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
			log.Fatal(err)
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
			log.Fatal(err)
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
			log.Fatal(err)
		}
	}

	fmt.Println("Data seeded successfully.")
}

func historicalPOSeed(db *sql.DB) {
	// Order data
	type historicalPO struct {
		CustomerID       string
		OrderDate        string
		SKUCode          string
		SKUID            string
		SKUName          string
		OrderQty         int
		OrderUnit        string
		UnitSellingPrice int
	}

	// Seed data
	seedData := []historicalPO{
		{"M1-PUMS-11", "20/12/2022", "SIK-8080", "SIK-080080-IBB", "Siku 80 mm x 80 mm x 6 mtr (KS/Ispat)", 10, "Batang", 581982},
		{"M1-PUMS-11", "20/12/2022", "SIK-9090", "SIK-090090-KS", "Siku 90 mm x 90 mm x 6 mtr (KS/Ispat)", 52, "Batang", 738739},
		{"M1-PUMS-11", "20/12/2022", "SIK-100100", "SIK-100100-IBB", "Siku 100 mm x 100 mm x 6 mtr (KS/Ispat)", 90, "Batang", 910360},
		{"M1-STQI-11", "17/01/2023", "PLT-BRDS23", "PLT-BRDS0230", "Besi Plat Bordes Tebal 2.3 mm x 1200 mm x 2400 mm", 100, "Lembar", 902832},
		{"M1-STQI-11", "17/01/2023", "PLT-SPHC1.55", "PLT-SPHC0155", "Plat Hitam Tebal 1.55 mm 1200 mm x 2400 mm", 200, "Lembar", 502978},
		{"M1-STQI-11", "17/01/2023", "PLT-SPHC1.8", "PLT-SPHC0180", "Plat Hitam Tebal 1.8 mm 1200 mm x 2400 mm", 50, "Lembar", 576828},
		{"M1-STQI-11", "19/01/2023", "PLT-SPHC2.6", "PLT-SPHC0026", "Plat Hitam Tebal 2.6 mm 1200 mm x 2400 mm", 10, "Lembar", 831644},
		{"M1-STQI-11", "19/01/2023", "PLT-SPHC2.8", "PLT-SPHC0280", "Plat Hitam Tebal 2.8 mm 1200 mm x 2400 mm", 20, "Lembar", 902927},
		{"M1-PUMS-11", "20/01/2023", "PIP-SCH404", "PIP-SCH404", "Pipa Hitam 4\" Sch x 6 m", 3, "Batang", 1557900},
		{"M1-PUMS-11", "20/01/2023", "PIP-SCH406", "PIP-SCH4060", "Pipa Hitam 6\" Sch x 6 m", 3, "Batang", 2724300},
		{"M1-PUMS-11", "20/01/2023", "PIP-SCH408", "PIP-SCH4080", "Pipa Hitam 8\" Sch x 6 m", 2, "Batang", 4158900},
		{"M1-PUMS-11", "20/01/2023", "PLT-SPHC4", "PLT-SPHC0400", "Besi Plat SPHC Tebal 6.0 mm x 1200 mm x 2400 mm (KS)", 3, "Lembar", 1782900},
		{"M1-PUMS-11", "20/01/2023", "PLT-SPHC5", "PLT-SPHC0500", "Besi Plat SPHC Tebal 8.0 mm x 1200 mm x 2400 mm (KS)", 5, "Lembar", 2364570},
		{"M1-PUMS-11", "20/01/2023", "PLT-SPHC6", "PLT-SPHC0600", "Besi Plat SPHC Tebal 10 mm x 1200 mm x 2400 mm (GG)", 4, "Lembar", 2946285},
		{"M1-PUMS-11", "20/01/2023", "PLT-SPHC10", "PLT-SPHC1000", "Besi Plat SPHC Tebal 16 mm x 1200 mm x 2400 mm (GG)", 1, "Lembar", 4766940},
		{"M1-PUMS-11", "20/01/2023", "PLT-SPHC8", "PLT-SPHC1200-GG", "Besi Plat SPHC Tebal 12 mm x 1200 mm x 2400 mm (GG)", 11, "Lembar", 3540600},
		{"M1-PUMS-11", "20/01/2023", "PLT-SPHC8", "PLT-SPHC1200-GG", "Besi Plat SPHC Tebal 12 mm x 1200 mm x 2400 mm (GG)", 14, "Lembar", 3540600},
		{"M1-PUMS-11", "20/01/2023", "SIK-4040", "SIK-040040-IBB", "Siku 40 mm x 40 mm x 6 mtr (KS)", 51, "Batang", 156600},
		{"M1-PUMS-11", "20/01/2023", "SIK-5050", "SIK-050050-IBB", "Siku 50 mm x 50 mm x 6 mtr (KS)", 72, "Batang", 246600},
		{"M1-PUMS-11", "20/01/2023", "SIK-6060", "SIK-060060-IBB", "Siku 60 mm x 60 mm x 6 mtr (KS)", 8, "Batang", 353700},
		{"M1-PUMS-11", "20/01/2023", "SIK-7070", "SIK-070070-IBB", "Siku 70 mm x 70 mm x 6 mtr (KS)", 8, "Batang", 482400},
		{"M1-PUMS-11", "20/01/2023", "SIK-8080", "SIK-080080-IBB", "Siku 80 mm x 80 mm x 6 mtr (KS)", 10, "Batang", 633150},
		{"M1-PUMS-11", "20/01/2023", "SIK-9090", "SIK-090090-KS", "Siku 90 mm x 90 mm x 6 mtr (KS)", 17, "Batang", 803700},
		{"M1-PUMS-11", "20/01/2023", "SIK-100100", "SIK-100100-IBB", "Siku 100 mm x 100 mm x 6 mtr (KS)", 1, "Batang", 990000},
		{"M1-PUMS-11", "20/01/2023", "SIK-100100", "SIK-100100-IBB", "Siku 100 mm x 100 mm x 6 mtr (KS)", 18, "Batang", 990000},
		{"M1-PUMS-11", "20/01/2023", "SIK-120120", "SIK-120120-IBB", "Siku 120 mm x 120 mm x 6 mtr (KS)", 5, "Batang", 1446300},
		{"M1-PUMS-11", "20/01/2023", "SIK-120120", "SIK-120120-IBB", "Siku 120 mm x 120 mm x 6 mtr (KS)", 20, "Batang", 1446300},
		{"M1-STQI-11", "20/01/2023", "PLT-SPHC2.8", "PLT-SPHC0280", "Plat Hitam Tebal 2.8 mm 1200 mm x 2400 mm", 20, "Lembar", 902927},
		{"M1-SUGP-11", "27/01/2023", "SIK-4040", "SIK-040040-IBB", "Besi Siku 40 mm x 40 mm x 4 mm x 6 m IBB", 30, "Batang", 162973},
		{"M1-SUGP-11", "27/01/2023", "SIK-5050", "SIK-050050-IBB", "Besi Siku 50 mm x 50 mm x 5 mm x 12 m IBB", 60, "Batang", 514147},
		{"M1-SUGP-11", "27/01/2023", "SIK-6060", "SIK-060060-IBB", "Besi Siku 60 mm x 60 mm x 6 mm x 12 m IBB", 50, "Batang", 737809},
		{"M1-SUGP-11", "27/01/2023", "SIK-7070", "SIK-070070-IBB", "Besi Siku 70 mm x 70 mm x 7 mm x 6 m IBB", 120, "Batang", 503237},
		{"M1-SUGP-11", "27/01/2023", "SIK-8080", "SIK-080080-IBB", "Besi Siku 80 mm x 80 mm x 8 mm x 12 m IBB", 15, "Batang", 1322680},
		{"M1-SUGP-11", "27/01/2023", "SIK-100100", "SIK-100100-IBB", "Besi Siku 100 mm x 100 mm x 8 mm x 12 m IBB", 30, "Batang", 1674972},
		{"M1-SUGP-11", "27/01/2023", "SIK-120120", "SIK-120120-IBB", "Besi Siku 120 mm x 120 mm x 10 mm x 12 m IBB", 11, "Batang", 2542785},
	}

	for _, data := range seedData {
		orderdate := formatDate(data.OrderDate)
		_, err := db.Exec(`
		INSERT INTO historypo (customer_id, order_date, sku_id, order_quantity, order_unit, unit_selling_price)
		VALUES ($1, $2, $3, $4, $5, $6)`,
			data.CustomerID, orderdate, data.SKUID, data.OrderQty, data.OrderUnit, data.UnitSellingPrice)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data seeded successfully.")

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
