CREATE TABLE IF NOT EXISTS subject (
    id INTEGER PRIMARY KEY,
    group_id INTEGER,
    class_id INTEGER,
    visiting_id INTEGER,
    grade_id INTEGER,
    FOREIGN KEY (group_id) REFERENCES "group"(id) ON DELETE CASCADE,
    FOREIGN KEY (class_id) REFERENCES class(id) ON DELETE CASCADE,
    FOREIGN KEY (visiting_id) REFERENCES visiting(id) ON DELETE CASCADE,
    FOREIGN KEY (grade_id) REFERENCES grade(id) ON DELETE CASCADE
);