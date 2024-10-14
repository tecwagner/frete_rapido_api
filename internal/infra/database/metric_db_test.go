package database_test

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/tecwagner/frete_rapido_api/internal/entities"
	"github.com/tecwagner/frete_rapido_api/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MetricDBTestSuite struct {
	suite.Suite
	db        *gorm.DB
	metricDB  *database.MetricDB
	carrierDB *database.CarrierDB
	carrier   *entities.Carrier
}

func (suite *MetricDBTestSuite) SetupSuite() {
	var err error

	suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Nil(err)

	err = suite.db.AutoMigrate(&entities.Carrier{})
	suite.Nil(err)

	suite.carrierDB = database.NewCarrierDB(suite.db)

	suite.metricDB = database.NewMetricDB(suite.db)
}

func TestMetricDBTestSuite(t *testing.T) {
	suite.Run(t, new(MetricDBTestSuite))
}

func (suite *MetricDBTestSuite) TestFind() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	carriers := []entities.Carrier{
		{
			Name:            "Carrier A",
			Service:         "Entrega Expressa",
			Deadline:        3,
			Price:           100.00,
			QuoteResponseID: 1,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			Name:            "Carrier B",
			Service:         "Entrega Normal",
			Deadline:        5,
			Price:           200.00,
			QuoteResponseID: 2,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	err := suite.carrierDB.Save(ctx, carriers, 1)
	suite.Nil(err)

	metricsResponse, err := suite.metricDB.Find(ctx, nil)
	suite.Nil(err)
	suite.NotNil(metricsResponse)

	sort.Slice(metricsResponse.CarrierMetrics, func(i, j int) bool {
		return metricsResponse.CarrierMetrics[i].CarrierName < metricsResponse.CarrierMetrics[j].CarrierName
	})

	suite.Equal(100.0, metricsResponse.CheapestFreight)
	suite.Equal(200.0, metricsResponse.MostExpensiveFreight)
	suite.Equal(2, len(metricsResponse.CarrierMetrics))

	suite.Equal("Carrier A", metricsResponse.CarrierMetrics[0].CarrierName)
	suite.Equal(1, metricsResponse.CarrierMetrics[0].Count)
	suite.Equal(100.0, metricsResponse.CarrierMetrics[0].TotalFreight)

	suite.Equal("Carrier B", metricsResponse.CarrierMetrics[1].CarrierName)
	suite.Equal(1, metricsResponse.CarrierMetrics[1].Count)
	suite.Equal(200.0, metricsResponse.CarrierMetrics[1].TotalFreight)
}
