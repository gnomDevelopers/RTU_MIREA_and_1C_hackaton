CREATE TABLE IF NOT EXISTS user_data (
    id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    last_name VARCHAR,
    first_name VARCHAR,
    father_name VARCHAR,
    university_id INTEGER REFERENCES university(id) ON DELETE CASCADE ,
    permission_id INTEGER DEFAULT 1,
    faculty_id INTEGER REFERENCES faculty(id) ON DELETE CASCADE,
    department_id INTEGER REFERENCES department(id) ON DELETE CASCADE,
    educational_direction VARCHAR,
    is_deleted BOOLEAN DEFAULT false,
    is_password_changed BOOLEAN DEFAULT false,
    created_at TIMESTAMP default now()
);