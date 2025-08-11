package postgres

import (
	"fmt"

	"github.com/Hot-One/monolith/config"
	role_model "github.com/Hot-One/monolith/models/role"
	user_model "github.com/Hot-One/monolith/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormConfig = gorm.Config

func ConnectPostgres(gormConfig *GormConfig, cfg config.Config) (*gorm.DB, error) {
	pgConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	db, err := gorm.Open(postgres.Open(pgConnStr), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		user_model.User{},
		role_model.Role{},
	)
}
