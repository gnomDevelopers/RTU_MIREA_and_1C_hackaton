CREATE TABLE IF NOT EXISTS "group" (
    id SERIAL PRIMARY KEY REFERENCES users(group_id),
    name VARCHAR
);
