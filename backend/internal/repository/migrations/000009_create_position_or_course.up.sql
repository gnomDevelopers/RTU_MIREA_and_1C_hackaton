CREATE TABLE IF NOT EXISTS position_or_course (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    user_data_id INTEGER REFERENCES user_data(id)
);