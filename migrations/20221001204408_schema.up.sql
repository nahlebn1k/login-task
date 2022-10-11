CREATE TABLE IF NOT EXISTS users
(
    id SERIAL CONSTRAINT users_pk PRIMARY KEY,
    login VARCHAR,
    password VARCHAR
);