CREATE TABLE IF NOT EXISTS visiting (
    id SERIAL PRIMARY KEY,
    class_id INTEGER REFERENCES class(id),
    user_id INTEGER REFERENCES users(id),
    type VARCHAR DEFAULT 'Н'
);