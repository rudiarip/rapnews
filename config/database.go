package config

import (
	"fmt"
	"rapnews/database/seeds"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) {

	dbConnString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Psql.User,
		cfg.Psql.Password,
		cfg.Psql.Host,
		cfg.Psql.Port,
		cfg.Psql.DBName,
	)

	// FIX UTAMA
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbConnString,
		PreferSimpleProtocol: true, // <— MATIKAN PREPARED STATEMENT
	}), &gorm.Config{
		PrepareStmt: false, // <— opsional, tapi bagus
	})

	if err != nil {
		log.Error().Err(err).Msg("[Connectionpostgres-1] Failed to connect to database " + cfg.Psql.Host)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[Connectionpostgres-2] Failed to connect to database")
		return nil, err
	}

	// Jalankan seeder
	seeds.SeedRoles(db)

	sqlDB.SetMaxOpenConns(cfg.Psql.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.Psql.DBMaxIdle)
	sqlDB.SetConnMaxLifetime(0)

	return &Postgres{
		DB: db,
	}, nil
}
