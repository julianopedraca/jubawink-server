package models

import (
	"context"
	"errors"
	"time"

	"github.com/julianopedraca/jubawink/internal/database"
)

const (
	Cycling = "cycling"
	Lifting = "lifting"
	Running = "running"
)

var validWorkoutTypes = map[string]bool{
	Cycling: true,
	Lifting: true,
	Running: true,
}

func isValidWorkoutType(workoutType string) bool {
	return validWorkoutTypes[workoutType]
}

type WorkoutSave struct {
	WorkoutType string `json:"workoutType" binding:"max=20"`
}

type WorkoutGetByUserId struct {
	UserId      int64     `json:"userId"`
	WorkoutType string    `json:"workoutType" binding:"max=20"`
	WorkoutDate time.Time `json:"workoutDate"`
}

func (ws *WorkoutSave) Save(userId int64) error {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if !isValidWorkoutType(ws.WorkoutType) {
		return errors.New("invalid workout type")
	}

	query := "INSERT INTO workouts(user_id, workout_type) VALUES ($1, $2)"
	_, err := db.Exec(ctx, query, userId, ws.WorkoutType)
	return err
}

func (w *WorkoutGetByUserId) GetWorkoutsByUserId() ([]WorkoutGetByUserId, error) {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `SELECT user_id, workout_type, workout_date FROM workouts WHERE user_id = $1`

	rows, err := db.Query(ctx, query, w.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []WorkoutGetByUserId

	for rows.Next() {
		var workout WorkoutGetByUserId
		if err := rows.Scan(&workout.UserId, &workout.WorkoutType, &workout.WorkoutDate); err != nil {
			return nil, err
		}
		workouts = append(workouts, workout)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workouts, nil
}
