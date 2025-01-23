DROP TABLE IF EXISTS LiftingWorkouts;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'workout_type') THEN
        DROP TYPE workout_type;
    END IF;
END $$;
