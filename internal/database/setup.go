package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

var Db *pgxpool.Pool

func ConnectDatabase() error {
	var err error
	Db, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Unable to create connection pool.")
		os.Exit(1)
	}

	log.Info("Database connected.")
	return nil
}
