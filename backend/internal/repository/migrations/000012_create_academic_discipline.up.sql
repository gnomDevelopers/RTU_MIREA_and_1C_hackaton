CREATE TABLE IF NOT EXISTS academic_discipline (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    educational_direction VARCHAR NOT NULL,
    semester INTEGER NOT NULL,
    lecture_hours INTEGER NOT NULL,
    practice_hours INTEGER NOT NULL,
    lab_hours INTEGER NOT NULL,
    individual_hours INTEGER NOT NULL,
    type_of_assessment VARCHAR NOT NULL
);