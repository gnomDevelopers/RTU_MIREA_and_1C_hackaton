CREATE TABLE IF NOT EXISTS gpa (
    user_id INTEGER REFERENCES users(id),
    value FLOAT NOT NULL
);