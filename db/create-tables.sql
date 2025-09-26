DROP TABLE IF EXISTS album;

CREATE TABLE album (
  id SERIAL PRIMARY KEY,
  title  VARCHAR(128) NOT NULL,
  artist VARCHAR(255) NOT NULL,
  price  NUMERIC(5,2) NOT NULL
);

INSERT INTO album (title, artist, price) VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);

-- CREATE DATABASE MAKE YOURSELF THE OWNER
-- sudo -u postgres psql
-- create database recordings;
-- alter database recordings owner to mugambi;

-- RUN THIS SQL SCRIPT
-- psql -d recordings
-- \i create-tables.sql