CREATE TABLE IF NOT EXISTS class (
    id INTEGER PRIMARY KEY,
    name VARCHAR,
    group_names VARCHAR[],
    teacher_names VARCHAR[],
    type VARCHAR,
    location VARCHAR,
    date VARCHAR,
    weekday INTEGER,
    week INTEGER,
    time_start VARCHAR,
    time_end VARCHAR
);