package models

import (
	"context"
	"time"

	"github.com/julianopedraca/jubawink/internal/database"
)

type User struct {
	UserName string `json:"username" binding:"required,max=50"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=6,max=21"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Save() error {
	db := database.Db
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := "INSERT INTO users(username, email, password_hash) VALUES ($1, $2, $3)"
	_, err := db.Query(ctx, query, u.UserName, u.Email, u.Password)
	return err
}

func (uc *UserCredentials) FindUserByEmail() (*UserCredentials, int64, error) {
	db := database.Db

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var creds UserCredentials
	var userId int64
	query := "SELECT email, password_hash, user_id FROM users WHERE email = $1"
	err := db.QueryRow(ctx, query, uc.Email).Scan(&creds.Email, &creds.Password, &userId)
	if err != nil {
		return nil, 0, err
	}
	return &creds, userId, err
}
