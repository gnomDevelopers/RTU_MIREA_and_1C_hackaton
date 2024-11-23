CREATE TABLE IF NOT EXISTS grade (
    id SERIAL PRIMARY KEY,
    class_id INTEGER REFERENCES class(id),
    user_id INTEGER REFERENCES users(id),
    value INTEGER NOT NULL
);