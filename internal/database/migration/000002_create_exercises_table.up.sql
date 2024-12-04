CREATE TABLE IF NOT EXISTS Exercises (
    exercise_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES Users(user_id) ON DELETE CASCADE,
    exercise_name VARCHAR(100) NOT NULL,
    sets INT NOT NULL CHECK (sets > 0),
    reps INT NOT NULL CHECK (reps > 0),
    weight NUMERIC(5, 2) CHECK (weight >= 0), -- weight in kg or lbs
    exercise_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP -- Automatically set to current date and time
);
