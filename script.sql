-- DROP TABLE categories;

CREATE TABLE IF NOT EXISTS categories (
	id uuid NOT NULL,
	"name" varchar(255) NOT NULL,
	CONSTRAINT categories_pkey PRIMARY KEY (id)
);

-- DROP TABLE services;

CREATE TABLE IF NOT EXISTS services (
	id uuid NOT NULL,
	description varchar(255) NOT NULL,
	price float8 NOT NULL,
	category_id uuid NOT NULL,
	CONSTRAINT service_pkey PRIMARY KEY (id)
);

-- DROP TABLE customers;

CREATE TABLE IF NOT EXISTS customers (
	id uuid NOT NULL,
	"name" varchar(255) NOT NULL,
	address varchar(255) NOT NULL,
	CONSTRAINT customer_pkey PRIMARY KEY (id)
);

-- DROP TABLE vehicles;

CREATE TABLE IF NOT EXISTS vehicles (
	id uuid NOT NULL,
	brand varchar(255) NOT NULL,
	model varchar(255) NOT NULL,
	"year" int4 NOT NULL,
	CONSTRAINT vehicles_pkey PRIMARY KEY (id)
);

-- DROP TABLE orders;

CREATE TABLE IF NOT EXISTS orders (
	id uuid NOT NULL,
	customer_id uuid NOT NULL,
	vehicle_id uuid NOT NULL,
	service_id uuid NOT NULL,
	request_date varchar(255) NOT NULL,
	status varchar(255) NOT NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (id)
);
