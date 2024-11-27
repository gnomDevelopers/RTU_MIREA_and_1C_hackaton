CREATE TABLE IF NOT EXISTS faculty (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    university VARCHAR
);

INSERT INTO faculty (name, university) VALUES ('null', 'null');