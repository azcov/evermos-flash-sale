package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // postgres driver
	"go.uber.org/zap"
)

type PGConnection struct {
	Host                  string
	Port                  int
	DBName                string
	User                  string
	Password              string
	SSLMode               string
	MaxOpenConnection     int
	MaxIdleConnection     int
	MaxConnectionLifetime time.Duration
}

// CreatePGConnection return db connection instance
func CreatePGConnection(config PGConnection) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s  sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		zap.S().Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		zap.S().Fatal(err)
	}

	// Setting database connection config
	db.SetMaxOpenConns(config.MaxOpenConnection)
	db.SetMaxIdleConns(config.MaxIdleConnection)
	db.SetConnMaxLifetime(config.MaxConnectionLifetime)

	zap.S().Info("Connected to PG DB Server: ", config.Host, " at port:", config.Port, " successfully!")

	return db, nil
}
