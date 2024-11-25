CREATE TABLE IF NOT EXISTS position_or_course (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    user_id INTEGER REFERENCES users(id)
);