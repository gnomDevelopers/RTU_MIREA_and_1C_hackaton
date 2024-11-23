CREATE TABLE IF NOT EXISTS my_schedule (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    date VARCHAR,
    time_start VARCHAR,
    time_end VARCHAR,
    user_data_id INTEGER REFERENCES user_data(id)
);