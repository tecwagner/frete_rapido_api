package config

import (
	"fmt"
	"log"

	"github.com/tecwagner/frete_rapido_api/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDatabase() *Database {
	return &Database{}
}

func NewTestDatabase() *gorm.DB {
	dbInstance := NewDatabase()
	dbInstance.setupTestConfig()

	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("test db error: %v", err)
	}

	return connection
}

func (d *Database) setupTestConfig() {
	d.Env = "test"
	d.DbTypeTest = "sqlite3"
	d.DsnTest = ":memory:"
	d.AutoMigrateDb = true
	d.Debug = true
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
	} else {
		d.Db, err = gorm.Open(sqlite.Open(d.DsnTest), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db = d.Db.Debug()
	}

	if d.AutoMigrateDb {
		err = d.autoMigrate()
		if err != nil {
			return nil, err
		}
	}

	return d.Db, nil
}

func (d *Database) autoMigrate() error {
	err := d.Db.AutoMigrate(&entities.Carrier{})
	if err != nil {
		return WrapError(err, "failed to auto-migrate database")
	}
	return nil
}

func WrapError(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}
