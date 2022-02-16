package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func setupTables(db *sql.DB) {
	if _, err := db.Exec(tableCreationQuery); err != nil {
			log.Fatal(err)
	}
}

const tableCreationQuery = `

DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS dogs CASCADE;
DROP TABLE IF EXISTS cats CASCADE;

CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);
CREATE TABLE dogs (
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name TEXT NOT NULL,
	is_good_boy BOOLEAN NOT NULL

CREATE TABLE IF NOT EXISTS cats
(
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name TEXT NOT NULL,
	weight FLOAT NOT NULL
);

`