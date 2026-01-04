package main

import (
	"embed"
	"errors"
	"fmt"
	config "go-clean-architecture-pzn/internal/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/gorm"
)

var fs embed.FS

func main() {
	viper := config.NewViper()
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error viper file: %w", err))
	// }

	log := config.NewLogger(viper)
	log.Info("Start migration")

	db := config.NewDatabase(viper, log)
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error database: %w", err))
	// }

	err := RunMigration(db)
	if err != nil {
		panic(fmt.Errorf("Error run migration: %w", err))
	}

	log.Info("Finish migration")
}

func RunMigration(db *gorm.DB) error {
	dbSql, err := db.DB()
	if err != nil {
		return err
	}

	location, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(dbSql, &mysql.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithInstance("iofs", location, "mysql", driver)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		} else {
			return err
		}
	}

	return nil
}
