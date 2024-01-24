-- migrations/1_init_schema.up.sql

CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    customer_id VARCHAR(12) NOT NULL UNIQUE,
    address VARCHAR(255) NOT NULL,
    city VARCHAR(35) NOT NULL,
    state VARCHAR(35) NOT NULL
);

CREATE UNIQUE INDEX idx_customer_id ON customers(customer_id);

CREATE TABLE IF NOT EXISTS suppliers (
    id SERIAL PRIMARY KEY,
    supplier_id VARCHAR(12) NOT NULL UNIQUE,
    address VARCHAR(255) NOT NULL,
    city VARCHAR(35) NOT NULL,
    state VARCHAR(35) NOT NULL
);

CREATE UNIQUE INDEX idx_supplier_id ON suppliers(supplier_id);

CREATE TABLE IF NOT EXISTS logistics (
    id SERIAL PRIMARY KEY,
    fleet_type VARCHAR(255) NOT NULL,
    capacity SERIAL NOT NULL,
    origin VARCHAR(35) NOT NULL,
    destination VARCHAR(35) NOT NULL,
    cost BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS rfqs (
    id SERIAL PRIMARY KEY,
    customer_id VARCHAR(12) NOT NULL,
    sku_id VARCHAR(20) NOT NULL,
    quantity SERIAL NOT NULL,
    unit_of_measure VARCHAR(12) /* its should be normalize to table units bit for now let it be strings */
);

CREATE TABLE IF NOT EXISTS pricelists (
    id SERIAL PRIMARY KEY,
    supplier_id VARCHAR(255) NOT NULL,
    sku_id VARCHAR(255) NOT NULL,
    price_per_unit FLOAT NOT NULL,
    stock_available INT NOT NULL /* stock - sku id may be will be separated from pricelist in the future */
);

CREATE TABLE IF NOT EXISTS historypo (
    id SERIAL PRIMARY KEY,
    customer_id VARCHAR(255) NOT NULL,
    order_date DATE NOT NULL,
    sku_id VARCHAR(20) NOT NULL,
    order_quantity INT NOT NULL,
    order_unit VARCHAR(12) NOT NULL,
    unit_selling_price FLOAT NOT NULL
);

/* all table still could be improved since this is just an mvp for price predictions we dont need that yet */