CREATE TABLE IF NOT EXISTS achievement (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    user_id INTEGER REFERENCES users(id)
);

