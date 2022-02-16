package goServer

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func SetupTables(db *sql.DB) {
	if _, err := db.Exec(tableCreationQuery); err != nil {
			log.Fatal(err)
	}
}

const tableCreationQuery = `

DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS cats CASCADE;

CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS cats
(
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name TEXT NOT NULL,
	weight FLOAT NOT NULL
);

INSERT INTO cats (name, weight)
VALUES ('kevin', 1.2), ('chungus', 42), ('pico', 0.001);
`