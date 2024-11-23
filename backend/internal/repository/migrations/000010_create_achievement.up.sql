CREATE TABLE IF NOT EXISTS achievement (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    user_data_id INTEGER REFERENCES user_data(id)
);

