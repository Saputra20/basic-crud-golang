package postgres

import (
	"basic-crud/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgConn *gorm.DB
var pgErr error

func init() {
	PGInit()
}

func PGInit() {
	cfg := config.Config()

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		cfg.PGHost,
		cfg.PGPort,
		cfg.PGDBName,
		cfg.PGUser,
		cfg.PGPassword,
	)

	pgConn, pgErr = gorm.Open(postgres.Open(dsn))
}

func PGConnection() (*gorm.DB, error) {
	if pgConn == nil || pgErr != nil {
		PGInit()
	}
	return pgConn, pgErr
}
