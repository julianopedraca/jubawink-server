package models

import (
	"context"
	"time"

	"github.com/julianopedraca/jubawink/internal/database"
)

type ExerciseLiftingSave struct {
	WorkoutId    int64  `json:"workoutId" `
	ExerciseName string `json:"exerciseName" binding:"max=100"`
	Sets         int64  `json:"sets"`
	Reps         int64  `json:"reps"`
	Weight       int64  `json:"weight"`
}

type ExerciseLiftingGet struct {
	UserId       int64  `json:"UserId"`
	WorkoutId    int64  `json:"workoutId" `
	ExerciseName string `json:"exerciseName" binding:"max=100"`
	Sets         int64  `json:"sets"`
	Reps         int64  `json:"reps"`
	Weight       int64  `json:"weight"`
}

type ExerciseCyclingSave struct {
	WorkoutId      int64 `json:"workoutId"`
	DistanceKm     int64 `json:"distanceKm"`
	AverageSpeed   int64 `json:"averageSpeed"`
	ElevationGainM int64 `json:"elevationGainM"`
	CaloriesBurned int64 `json:"caloriesBurned"`
}

type ExerciseCyclingGet struct {
	UserId         int64 `json:"userId"`
	WorkoutId      int64 `json:"workoutId"`
	DistanceKm     int64 `json:"distanceKm"`
	AverageSpeed   int64 `json:"averageSpeed"`
	ElevationGainM int64 `json:"elevationGainM"`
	CaloriesBurned int64 `json:"caloriesBurned"`
}

type ExerciseRunningSave struct {
	WorkoutId      int64  `json:"workoutId"`
	DistanceKm     int64  `json:"distanceKm"`
	AveragePace    string `json:"averagePace"`
	CaloriesBurned int64  `json:"caloriesBurned"`
}

type ExerciseRunningGet struct {
	UserId         int64   `json:"userId"`
	DistanceKm     float64 `json:"distanceKm"`
	AveragePace    string  `json:"averagePace"`
	CaloriesBurned int     `json:"caloriesBurned"`
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
		WHERE user_id = ($1)
	`

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

func (ecs *ExerciseCyclingSave) SaveCycling() error {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := "INSERT INTO CyclingWorkouts (workout_id, distance_km, average_speed, elevation_gain_m, calories_burned) VALUES ($1, $2, $3, $4, $5)"

	_, err := db.Query(ctx, query, ecs.WorkoutId, ecs.DistanceKm, ecs.AverageSpeed, ecs.ElevationGainM, ecs.CaloriesBurned)
	return err
}

func (ecg *ExerciseCyclingGet) GetCycling() ([]ExerciseCyclingGet, error) {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
		SELECT 
			workouts.user_id, 
			cyclingworkouts.distance_km, 
			cyclingworkouts.average_speed, 
			cyclingworkouts.elevation_gain_m, 
			cyclingworkouts.calories_burned 
		FROM cyclingworkouts 
		INNER JOIN workouts 
		ON cyclingworkouts.workout_id = workouts.workout_id
		WHERE user_id = ($1)
	`

	rows, err := db.Query(ctx, query, ecg.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercisesCycling []ExerciseCyclingGet

	for rows.Next() {
		var exerciseCycling ExerciseCyclingGet
		if err := rows.Scan(&exerciseCycling.UserId, &exerciseCycling.DistanceKm, &exerciseCycling.AverageSpeed, &exerciseCycling.ElevationGainM, &exerciseCycling.CaloriesBurned); err != nil {
			return nil, err
		}
		exercisesCycling = append(exercisesCycling, exerciseCycling)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exercisesCycling, nil
}

func (ers *ExerciseRunningSave) SaveRunning() error {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
        INSERT INTO RunningWorkouts (workout_id, distance_km, average_pace, calories_burned)
        VALUES ($1, $2, $3, $4)
    `

	_, err := db.Exec(ctx, query, ers.WorkoutId, ers.DistanceKm, ers.AveragePace, ers.CaloriesBurned)
	if err != nil {
		return err
	}

	return nil
}

func (erg *ExerciseRunningGet) GetRunning() ([]ExerciseRunningGet, error) {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
        SELECT 
            workouts.user_id, 
            runningworkouts.distance_km, 
            runningworkouts.average_pace, 
            runningworkouts.calories_burned 
        FROM runningworkouts 
        INNER JOIN workouts 
        ON runningworkouts.workout_id = workouts.workout_id
        WHERE user_id = ($1)
    `

	rows, err := db.Query(ctx, query, erg.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercisesRunning []ExerciseRunningGet

	for rows.Next() {
		var exerciseRunning ExerciseRunningGet
		if err := rows.Scan(&exerciseRunning.UserId, &exerciseRunning.DistanceKm, &exerciseRunning.AveragePace, &exerciseRunning.CaloriesBurned); err != nil {
			return nil, err
		}
		exercisesRunning = append(exercisesRunning, exerciseRunning)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return exercisesRunning, nil
}
