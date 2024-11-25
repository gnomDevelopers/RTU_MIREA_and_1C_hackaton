CREATE TABLE IF NOT EXISTS faculty (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    university VARCHAR
);

--INSERT INTO faculty (id, name, university) VALUES (-1, 'admin', 'admin');