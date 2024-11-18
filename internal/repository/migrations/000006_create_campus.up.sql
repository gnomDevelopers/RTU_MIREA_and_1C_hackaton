CREATE TABLE IF NOT EXISTS campus (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    university_id INTEGER REFERENCES university(id),
    address VARCHAR
);