CREATE TABLE IF NOT EXISTS RunningWorkouts (
    running_id SERIAL PRIMARY KEY,
    workout_id INTEGER NOT NULL,
    distance_km DECIMAL(5, 2) NOT NULL, -- Distância percorrida em km
    average_pace VARCHAR(20), -- Ritmo médio (ex: "5:30 min/km")
    calories_burned INTEGER, -- Calorias queimadas
    CONSTRAINT fk_workout
        FOREIGN KEY (workout_id)
        REFERENCES Workouts (workout_id)
        ON DELETE CASCADE
);
