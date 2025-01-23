CREATE TABLE IF NOT EXISTS CyclingWorkouts (
    cycling_id SERIAL PRIMARY KEY,
    workout_id INTEGER NOT NULL,
    distance_km DECIMAL(5, 2) NOT NULL, -- Distância percorrida em km
    average_speed DECIMAL(5, 2), -- Velocidade média em km/h
    elevation_gain_m INTEGER, -- Ganho de elevação em metros
    calories_burned INTEGER, -- Calorias queimadas
    CONSTRAINT fk_workout
        FOREIGN KEY (workout_id)
        REFERENCES Workouts (workout_id)
        ON DELETE CASCADE
);
