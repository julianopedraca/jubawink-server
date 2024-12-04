package models

import (
	"context"
	"time"

	"github.com/julianopedraca/jubawink/internal/database"
)

type Exercise struct {
	UserId       int64  `json:"userId" swaggerignore:"true"`
	ExerciseName string `json:"exerciseName" binding:"max=100"`
	Sets         int64  `json:"sets"`
	Reps         int64  `json:"reps"`
	Weight       int64  `json:"weight"`
}

func (e *Exercise) Save() error {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := "INSERT INTO exercises (user_id, exercise_name, sets, reps, weight) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Query(ctx, query, e.UserId, e.ExerciseName, e.Sets, e.Reps, e.Weight)
	return err
}
