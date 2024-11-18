CREATE TABLE IF NOT EXISTS "group" (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    user_id INTEGER REFERENCES "user"(id)
);
