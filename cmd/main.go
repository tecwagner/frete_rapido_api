package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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

	if err := loadEnvVars(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	dbConnection, err := setupDatabase()
	if err != nil {
		log.Fatalf("Error starting the database: %v", err)
	}

	freightFastAPIURL := os.Getenv("API_FRETE_RAPIDO_URL")
	if freightFastAPIURL == "" {
		log.Fatal("API_FRETE_RAPIDO_URL is not set")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	quoteService := createquote.NewQuoteService(freightFastAPIURL, client)

	carrierDB := database.NewCarrierDB(dbConnection)

	fetcherFunc := func(ctx context.Context, request createquote.CreateQuoteInputDTO) (createquote.FreightFastOutputDTO, error) {
		return quoteService.GetQuoteFromFreightFast(ctx, request)
	}

	createquoteUseCase := createquote.NewCreateQuoteUseCase(carrierDB, fetcherFunc)

	webserverInstance := webserver.NewWebServer(":8081")

	webQuoteHandler := web.NewWebQuoteHandler(*createquoteUseCase)
	webserverInstance.AddHandlerPublic("POST", "/api/v1/quote", webQuoteHandler.CreateQuote)

	fmt.Println("Server is running", webserverInstance.WebServerPort)
	fmt.Println("Welcome to Frete Rápido API")
	log.Info("Starting Frete Rápido Application")

	webserverInstance.Start()

}

func setupDatabase() (*gorm.DB, error) {
	autoMigrateDb, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE_DB"))
	if err != nil {
		return nil, fmt.Errorf("error parsing boolean from AUTO_MIGRATE_DB: %v", err)
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return nil, fmt.Errorf("error parsing boolean from DEBUG: %v", err)
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
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return dbConnection, nil
}

func loadEnvVars() error {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file: ", err)
		return err
	}
	return nil
}
