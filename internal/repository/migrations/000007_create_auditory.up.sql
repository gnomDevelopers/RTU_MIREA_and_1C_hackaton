CREATE TABLE IF NOT EXISTS auditory (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    campus_id INTEGER REFERENCES campus(id),
    type VARCHAR,
    profile VARCHAR,
    capacity INTEGER
);