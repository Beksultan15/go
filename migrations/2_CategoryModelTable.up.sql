CREATE TABLE category (
     id bigserial PRIMARY KEY,
    category_name varchar(255) not null
    
    );

INSERT INTO category (category_name) VALUES ('Food');
INSERT INTO category (category_name) VALUES ('Digital');
INSERT INTO category (category_name) VALUES ('Furniture');
INSERT INTO category (category_name) VALUES ('Health');
INSERT INTO category (category_name) VALUES ('Other');