CREATE TABLE IF NOT EXISTS subject (
    id SERIAL PRIMARY KEY,
    group_id INTEGER REFERENCES "group"(id),
    class_id INTEGER REFERENCES class(id),
    visiting_id INTEGER REFERENCES visiting(id),
    grade_id INTEGER REFERENCES grade(id)
);