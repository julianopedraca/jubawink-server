package models

import (
	"context"
	"time"

	"github.com/julianopedraca/jubawink/internal/database"
)

type ExerciseLiftingSave struct {
	WorkoutId    int64  `json:"WorkoutId" `
	ExerciseName string `json:"exerciseName" binding:"max=100"`
	Sets         int64  `json:"sets"`
	Reps         int64  `json:"reps"`
	Weight       int64  `json:"weight"`
}

type ExerciseLiftingGet struct {
	UserId       int64  `json:"UserId"`
	WorkoutId    int64  `json:"WorkoutId" `
	ExerciseName string `json:"exerciseName" binding:"max=100"`
	Sets         int64  `json:"sets"`
	Reps         int64  `json:"reps"`
	Weight       int64  `json:"weight"`
}

func (els *ExerciseLiftingSave) SaveLifting() error {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := "INSERT INTO LiftingWorkouts (workout_id, exercise_name, weight_kg, repetitions, sets) VALUES ($1, $2, $3, $4, $5)"

	_, err := db.Query(ctx, query, els.WorkoutId, els.ExerciseName, els.Weight, els.Reps, els.Sets)
	return err
}

func (elg *ExerciseLiftingGet) GetLifting() ([]ExerciseLiftingGet, error) {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
	SELECT 
		workouts.user_id, 
		liftingworkouts.workout_id, 
		liftingworkouts.exercise_name, 
		liftingworkouts.weight_kg, 
		liftingworkouts.repetitions, 
		liftingworkouts.sets
	FROM liftingworkouts 
	INNER JOIN workouts 
	ON liftingworkouts.workout_id = workouts.workout_id
	WHERE user_id = ($1)`

	rows, err := db.Query(ctx, query, elg.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercisesLifting []ExerciseLiftingGet

	for rows.Next() {
		var exerciseLifting ExerciseLiftingGet
		if err := rows.Scan(&exerciseLifting.UserId, &exerciseLifting.WorkoutId, &exerciseLifting.ExerciseName, &exerciseLifting.Weight, &exerciseLifting.Reps, &exerciseLifting.Sets); err != nil {
			return nil, err
		}
		exercisesLifting = append(exercisesLifting, exerciseLifting)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exercisesLifting, nil
}
