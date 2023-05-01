CREATE TABLE users (
    id bigserial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    username      varchar(255) not null unique,
    password varchar(255) not null,
    email         varchar(255) not null
);

INSERT INTO users(
	id, first_name,last_name,username,password,email)
	VALUES (1, 'Test','test','test','qwerty1','test@mail.ru');
