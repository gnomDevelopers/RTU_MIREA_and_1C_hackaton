CREATE TABLE IF NOT EXISTS role (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    user_data_id INTEGER REFERENCES user_data(id)
);