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

func (els *ExerciseLiftingSave) Save() error {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := "INSERT INTO LiftingWorkouts (workout_id, exercise_name, weight_kg, repetitions, sets) VALUES ($1, $2, $3, $4, $5)"

	_, err := db.Query(ctx, query, els.WorkoutId, els.ExerciseName, els.Weight, els.Reps, els.Sets)
	return err
}
