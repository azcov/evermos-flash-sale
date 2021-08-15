package init

import (
	"database/sql"
	"os"

	"github.com/azcov/evermos-flash-sale/internal/db"
)

// ConnectToPGServer is a function to init PostgreSQL connection
func ConnectToPGServer(config *Config) (*sql.DB, error) {

	dbpg, err := db.CreatePGConnection(db.PGConnection{
		Host:                  config.Database.Pg.Host,
		Port:                  config.Database.Pg.Port,
		User:                  config.Database.Pg.User,
		Password:              config.Database.Pg.Password,
		DBName:                config.Database.Pg.DBName,
		SSLMode:               config.Database.Pg.SSLMode,
		MaxOpenConnection:     config.Database.Pg.MaxOpenConnection,
		MaxIdleConnection:     config.Database.Pg.MaxIdleConnection,
		MaxConnectionLifetime: config.Database.Pg.MaxConnectionLifetime,
	})

	if err != nil {
		os.Exit(1)
	}

	return dbpg, err
}
