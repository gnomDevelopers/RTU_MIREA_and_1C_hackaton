CREATE TABLE IF NOT EXISTS user_data (
    id INTEGER PRIMARY KEY REFERENCES "user"(id),
    last_name VARCHAR,
    first_name VARCHAR,
    father_name VARCHAR,
    university_id INTEGER REFERENCES university(id),
    permission_id INTEGER,
    faculty_id INTEGER REFERENCES faculty(id),
    department_id INTEGER REFERENCES department(id),
    educational_direction VARCHAR,
    is_deleted BOOLEAN DEFAULT false,
    is_password_changed BOOLEAN DEFAULT false,
    created_at TIMESTAMP default now()
);