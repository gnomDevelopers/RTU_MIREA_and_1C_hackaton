CREATE TABLE IF NOT EXISTS academic_discipline (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    educational_direction_id INTEGER REFERENCES educational_direction(id),
    semester INTEGER,
    lecture_hours INTEGER,
    practice_hours INTEGER,
    lab_hours INTEGER,
    individual_hours INTEGER,
    type_of_assessment VARCHAR
);