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

// Construtor para banco de dados principal
func NewDb() *Database {
	return &Database{}
}

// Construtor para banco de dados de testes
func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	// Conexão com banco de dados de teste
	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("test db error: %v", err)
	}

	return connection
}

// Método de conexão com o banco de dados
func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "test" {
		// Conexão com banco de dados de produção
		d.Db, err = gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
	} else {
		// Conexão com banco de dados de teste (SQLite em memória)
		d.Db, err = gorm.Open(sqlite.Open(d.DsnTest), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		// Habilitar modo debug para mostrar logs SQL
		d.Db = d.Db.Debug()
	}

	// Executar migrações automáticas, se habilitado
	if d.AutoMigrateDb {
		err = d.Db.AutoMigrate(&entities.Carrier{})
		if err != nil {
			return nil, err
		}
	}

	// Retorna a conexão com o banco de dados
	return d.Db, nil
}

// Função de erro formatado
func wrapError(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

// CarrierDB usando Gorm
type CarrierDB struct {
	DB *gorm.DB
}

func NewCarrierDB(db *gorm.DB) *CarrierDB {
	return &CarrierDB{
		DB: db,
	}
}

// Método para salvar carrier usando Gorm
func (c *CarrierDB) Save(ctx context.Context, carriers []entities.Carrier, quoteResponseID uint) error {

	// Inicia uma transação
	tx := c.DB.WithContext(ctx).Begin()

	// Verifica se a transação foi iniciada corretamente
	if tx.Error != nil {
		return wrapError(tx.Error, "failed to begin transaction")
	}

	// Itera sobre cada carrier e associa o quoteResponseID
	for i := range carriers {
		carriers[i].QuoteResponseID = quoteResponseID // Associa o carrier à quote

		// Tenta salvar o carrier na transação
		if err := tx.Create(&carriers[i]).Error; err != nil {
			// Reverte a transação em caso de erro
			tx.Rollback()
			return wrapError(err, "failed to save carrier")
		}
	}

	// Commit a transação se tudo der certo
	if err := tx.Commit().Error; err != nil {
		return wrapError(err, "failed to commit transaction")
	}

	return nil
}
