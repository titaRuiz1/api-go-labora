-- psql
CREATE DATABASE labora_proyect_1;

-- \c labora_proyect_1

CREATE TABLE items (
   id SERIAL PRIMARY KEY,
   customer_name VARCHAR(100) NOT NULL,
   order_date DATE NOT NULL,
   product VARCHAR(255) NOT NULL,
   quantity INT NOT NULL,
   price NUMERIC NOT NULL
);