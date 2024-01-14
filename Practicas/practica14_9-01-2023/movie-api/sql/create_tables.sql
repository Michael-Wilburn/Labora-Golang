CREATE TABLE movie (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created DATE,
    stock INTEGER DEFAULT 10,
    price DECIMAL(10, 2) NOT NULL
);