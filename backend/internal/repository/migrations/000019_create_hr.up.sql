CREATE TABLE IF NOT EXISTS hr (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR,
    first_name VARCHAR,
    father_name VARCHAR,
    company VARCHAR,
    image_link TEXT
);