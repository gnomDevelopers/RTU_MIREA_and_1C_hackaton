CREATE TABLE IF NOT EXISTS department (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    university VARCHAR
);

INSERT INTO department (name) VALUES ('null');