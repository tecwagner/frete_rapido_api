package main

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/tecwagner/frete_rapido_api/internal/infra/database"
	createquote "github.com/tecwagner/frete_rapido_api/internal/useCase/create_quote"
	"github.com/tecwagner/frete_rapido_api/web"
	webserver "github.com/tecwagner/frete_rapido_api/web/webserver"
	"gorm.io/gorm"
)

func main() {

	dbConnection, err := setupDatabase()
	if err != nil {
		log.Fatalf("Erro ao iniciar o banco de dados: %v", err)
	}

	carrierDB := database.NewCarrierDB(dbConnection)
	fetcherFunc := createquote.GetQuoteFromFreightFast

	createquoteUseCase := createquote.NewCreateQuoteUseCase(carrierDB, fetcherFunc)

	webserverInstance := webserver.NewWebServer(":8081")

	webQuoteHandler := web.NewWebQuoteHandler(*createquoteUseCase)

	webserverInstance.AddHandlerPublic("POST", "/api/v1/quote", webQuoteHandler.CreateQuote)

	fmt.Println("Server is running", webserverInstance.WebServerPort)
	fmt.Println("Seja Bem-Vindo ao Frete Rápido API")
	log.Info("Iniciando Aplicação de Frete Rápido")

	webserverInstance.Start()

}

func setupDatabase() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar o arquivo .env: %v", err)
	}

	autoMigrateDb, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE_DB"))
	if err != nil {
		return nil, fmt.Errorf("erro ao analisar booleano de AUTO_MIGRATE_DB: %v", err)
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return nil, fmt.Errorf("erro ao analisar booleano de DEBUG: %v", err)
	}

	db := &database.Database{
		AutoMigrateDb: autoMigrateDb,
		Debug:         debug,
		DsnTest:       os.Getenv("DSN_TEST"),
		Dsn:           os.Getenv("DSN"),
		DbTypeTest:    os.Getenv("DB_TYPE_TEST"),
		DbType:        os.Getenv("DB_TYPE"),
		Env:           os.Getenv("ENV"),
	}

	dbConnection, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	return dbConnection, nil
}
