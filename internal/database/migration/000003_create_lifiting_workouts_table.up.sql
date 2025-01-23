CREATE TABLE IF NOT EXISTS LiftingWorkouts (
    lifting_id SERIAL PRIMARY KEY,
    workout_id INTEGER NOT NULL,
    exercise_name VARCHAR(50) NOT NULL, -- Nome do exercício (ex: "Agachamento", "Supino")
    weight_kg DECIMAL(5, 2) NOT NULL, -- Peso utilizado em kg
    repetitions INTEGER NOT NULL, -- Número de repetições
    sets INTEGER NOT NULL, -- Número de séries
    CONSTRAINT fk_workout
        FOREIGN KEY (workout_id)
        REFERENCES Workouts (workout_id)
        ON DELETE CASCADE
);
