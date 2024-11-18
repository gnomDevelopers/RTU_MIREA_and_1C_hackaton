CREATE TABLE IF NOT EXISTS class (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    academic_discipline_id INTEGER REFERENCES academic_discipline(id),
    group_names VARCHAR[],
    teacher_names VARCHAR[],
    type VARCHAR,
    auditory_id INTEGER REFERENCES auditory(id),
    date VARCHAR,
    weekday INTEGER,
    week INTEGER,
    time_start VARCHAR,
    time_end VARCHAR
);