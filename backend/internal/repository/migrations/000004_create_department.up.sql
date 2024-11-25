CREATE TABLE IF NOT EXISTS department (
    id SERIAL PRIMARY KEY,
    name VARCHAR
);

INSERT INTO department (name) VALUES ('admin');