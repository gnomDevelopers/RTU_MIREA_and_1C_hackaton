CREATE TABLE IF NOT EXISTS achievement (
    id SERIAL PRIMARY KEY,
    name TEXT,
    result VARCHAR,
    user_id INTEGER REFERENCES users(id)
);

