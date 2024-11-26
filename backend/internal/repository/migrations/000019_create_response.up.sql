CREATE TABLE IF NOT EXISTS response (
     id SERIAL PRIMARY KEY,
     hr_id INTEGER REFERENCES hr(id),
     work_user_id INTEGER REFERENCES work_user(id)
);