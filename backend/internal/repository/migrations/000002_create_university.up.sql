CREATE TABLE IF NOT EXISTS university (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    postfix VARCHAR
);

INSERT INTO university (name, postfix) VALUES ('null', 'null');