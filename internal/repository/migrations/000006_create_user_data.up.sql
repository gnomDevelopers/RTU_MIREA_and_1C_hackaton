CREATE TABLE IF NOT EXISTS user_data (
    id INTEGER PRIMARY KEY,
    last_name VARCHAR,
    first_name VARCHAR,
    father_name VARCHAR,
    role_id INTEGER,
    achievement_id INTEGER,
    faculty_id INTEGER,
    department_id INTEGER,
    position_or_course VARCHAR,
    is_deleted BOOLEAN,
    created_at TIMESTAMP,
    FOREIGN KEY (id) REFERENCES "user"(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    FOREIGN KEY (achievement_id) REFERENCES achievement(id) ON DELETE CASCADE,
    FOREIGN KEY (faculty_id) REFERENCES faculty(id) ON DELETE CASCADE,
    FOREIGN KEY (department_id) REFERENCES department(id) ON DELETE CASCADE
);
