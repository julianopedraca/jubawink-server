package models

import (
	"context"
	"errors"
	"time"

	"github.com/julianopedraca/jubawink/internal/database"
)

type Workout struct {
	UserId      int64  `json:"userId" binding:"required"`
	WorkoutType string `json:"workoutType" binding:"max=20"`
}

func (w *Workout) Save() error {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if w.WorkoutType != "Cycling" && w.WorkoutType != "Lifting" && w.WorkoutType != "Running" {
		return errors.New("invalid workout type")
	}

	query := "INSERT INTO workouts(user_id, workout_type) VALUES ($1, $2)"
	_, err := db.Query(ctx, query, w.UserId, w.WorkoutType)
	return err
}

func (w *Workout) GetWorkoutsByUserId(userId int64) ([]Workout, error) {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `SELECT workouts.* FROM workouts INNER JOIN users ON workouts.user_id = users.user_id WHERE workouts.user_id = $1`

	rows, err := db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []Workout

	for rows.Next() {
		var workout Workout
		if err := rows.Scan(&workout.UserId, &workout.WorkoutType); err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workouts, nil
}
