

CREATE TABLE IF NOT EXISTS class (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    group_names VARCHAR[],
    teacher_names VARCHAR[],
    type VARCHAR,
    auditory VARCHAR,
    date VARCHAR,
    weekday INTEGER,
    week INTEGER,
    time_start VARCHAR,
    time_end VARCHAR,
    university VARCHAR
);