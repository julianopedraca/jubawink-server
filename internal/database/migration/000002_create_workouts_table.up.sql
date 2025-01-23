DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'workout_type') THEN
        CREATE TYPE workout_type AS ENUM ('lifting', 'running', 'cycling');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS Workouts (
    workout_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    workout_type workout_type NOT NULL, -- Tipo de treino: "Corrida", "Levantamento", "Ciclismo"
    workout_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP, -- Data do treino
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES Users (user_id)
        ON DELETE CASCADE
);
