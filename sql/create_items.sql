-- psql
CREATE DATABASE labora_proyect_3;

-- \c labora_proyect_3

CREATE TABLE items (
   id SERIAL PRIMARY KEY,
   customer_name VARCHAR(100) NOT NULL,
   order_date DATE NOT NULL,
   product VARCHAR(255) NOT NULL,
   quantity INT NOT NULL,
   price NUMERIC NOT NULL
);
INSERT INTO items (customer_name, order_date, product, quantity, price)
VALUES ('Jesus', '2023-05-07', 'Coca-Cola', 2, 1.5),
       ('Jorge', '2023-05-08', 'Pepsi', 3, 2.0),
       ('Luis', '2023-05-08', 'Sprite', 4, 2.5),
       ('Ximena', '2023-05-09', 'Fanta', 1, 1.0),
       ('Sara', '2023-05-09', 'Agua Mineral', 6, 1.0);

-- SELECT *FROM items;    


