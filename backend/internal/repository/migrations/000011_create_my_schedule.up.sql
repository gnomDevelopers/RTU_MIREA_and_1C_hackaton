CREATE TABLE IF NOT EXISTS my_schedule (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    date VARCHAR,
    time_start VARCHAR,
    time_end VARCHAR,
    user_id INTEGER REFERENCES users(id)
);