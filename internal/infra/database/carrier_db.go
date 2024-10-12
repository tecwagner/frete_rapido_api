package database

import (
	"context"
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

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("test db error: %v", err)
	}

	return connection
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
		err = d.Db.AutoMigrate(&entities.Carrier{})
		if err != nil {
			return nil, err
		}
	}

	return d.Db, nil
}

func wrapError(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

type CarrierDB struct {
	DB *gorm.DB
}

func NewCarrierDB(db *gorm.DB) *CarrierDB {
	return &CarrierDB{
		DB: db,
	}
}

func (c *CarrierDB) Save(ctx context.Context, carriers []entities.Carrier, quoteResponseID uint) error {

	tx := c.DB.WithContext(ctx).Begin()

	if tx.Error != nil {
		return wrapError(tx.Error, "failed to begin transaction")
	}

	for i := range carriers {
		carriers[i].QuoteResponseID = quoteResponseID

		if err := tx.Create(&carriers[i]).Error; err != nil {
			tx.Rollback()
			return wrapError(err, "failed to save carrier")
		}
	}

	if err := tx.Commit().Error; err != nil {
		return wrapError(err, "failed to commit transaction")
	}

	return nil
}
