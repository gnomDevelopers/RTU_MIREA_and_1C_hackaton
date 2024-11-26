CREATE TABLE IF NOT EXISTS work_user (
    id SERIAL PRIMARY KEY,
    phone_number varchar,
    telegram varchar,
    skills varchar[],
    cv_path varchar,
    hide_profile varchar
);