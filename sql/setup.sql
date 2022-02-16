-- Use this file to define your SQL tables
-- The SQL in this file will be executed when you run `npm run setup-db`

DROP TABLE IF EXISTS dogs;
DROP TABLE IF EXISTS cats;
DROP TABLE IF EXISTS planes;
DROP TABLE IF EXISTS trains;
DROP TABLE IF EXISTS automobiles;

CREATE TABLE dogs (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name TEXT NOT NULL,
  is_good_boy BOOLEAN NOT NULL
);

CREATE TABLE cats (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name TEXT NOT NULL,
  weight FLOAT NOT NULL
);

CREATE TABLE planes (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  model TEXT NOT NULL,
  cost MONEY NOT NULL
);

CREATE TABLE trains (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  model TEXT NOT NULL,
  manufacturer TEXT NOT NULL
);

CREATE TABLE automobiles (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  model TEXT NOT NULL,
  hp DOUBLE PRECISION NOT NULL
);

INSERT INTO dogs (name, is_good_boy)
VALUES ('spot', true), ('jeep', true), ('jeff', true);

INSERT INTO cats (name, weight)
VALUES ('kevin', 1.2), ('chungus', 42), ('pico', 0.001);

INSERT INTO planes (model, cost)
VALUES ('mustang', '$3.50'), ('777', '$123,456.12'), ('787', '$7,654,321.01');

INSERT INTO trains (model, manufacturer)
VALUES ('train1', 'train co'), ('train2', 'train co'), ('train3', 'train co');

INSERT INTO automobiles (model, hp)
VALUES ('mustang', 350.789654321), ('bug', 63.2), ('transit', 213.111);