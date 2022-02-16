package goServer

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

func SetupTables(db *sql.DB) {
	if _, err := db.Exec(tableCreationQuery); err != nil {
			log.Fatal(err)
	}
}

const tableCreationQuery = `

DROP TABLE IF EXISTS products CASCADE;

CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)

`