CREATE TABLE IF NOT EXISTS work_user (
    id SERIAL PRIMARY KEY,
    speciality VARCHAR,
    work_experience VARCHAR,
    additional_experience VARCHAR,
    useful_links VARCHAR[],
    phone_number VARCHAR,
    telegram VARCHAR,
    skills VARCHAR[],
    cv_path VARCHAR,
    hide_profile VARCHAR
);