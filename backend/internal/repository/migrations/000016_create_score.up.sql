CREATE TABLE IF NOT EXISTS score (
    user_id INTEGER REFERENCES users(id),
    sum INTEGER NOT NULL,
    count INTEGER NOT NULL,
    subject_name VARCHAR NOT NULL
);