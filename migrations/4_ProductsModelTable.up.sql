

CREATE TABLE products ( 
    id bigserial PRIMARY KEY,
    category_id bigint NOT NULL REFERENCES category ON DELETE CASCADE,
    product_title varchar(255) not null,
    price int not null,
    quantity int not null,
    product_description varchar(255) not null);


INSERT INTO products(
	id, product_title ,category_id,price, quantity, product_description)
	VALUES (1, 'Test', 1,123, 1, 'Test product');