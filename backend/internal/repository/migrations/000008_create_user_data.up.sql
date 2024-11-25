CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR,
    password VARCHAR,
    last_name VARCHAR,
    first_name VARCHAR,
    father_name VARCHAR,
    university_id INTEGER REFERENCES university(id) ON DELETE CASCADE ,
    role VARCHAR NOT NULL,
    group_id INTEGER REFERENCES "group"(id) ON DELETE CASCADE ,
    faculty_id INTEGER REFERENCES faculty(id) ON DELETE CASCADE,
    department_id INTEGER REFERENCES department(id) ON DELETE CASCADE,
    educational_direction VARCHAR,
    is_deleted BOOLEAN DEFAULT false,
    is_password_changed BOOLEAN DEFAULT false,
    created_at TIMESTAMP default now()
);