CREATE TABLE users (
    id bigserial not null PRIMARY key,
    email VARCHAR not null UNIQUE,
    encrypted_password VARCHAR not NULL
);