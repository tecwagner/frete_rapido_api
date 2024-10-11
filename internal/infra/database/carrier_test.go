package database

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/tecwagner/frete_rapido_api/internal/entities"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

type CarrierDBTestSuite struct {
	suite.Suite
	db        *gorm.DB
	carrierDB *CarrierDB
	carrier   *entities.Carrier
}

func (suite *CarrierDBTestSuite) SetupSuite() {
	var err error

	suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	suite.Nil(err)

	// Executa a migration para criar a tabela "carriers"
	err = suite.db.AutoMigrate(&entities.Carrier{})
	suite.Nil(err)

	// Cria uma nova instância do CarrierDB utilizando o Gorm
	suite.carrierDB = NewCarrierDB(suite.db)
}

func TestCarrierDBTestSuite(t *testing.T) {
	suite.Run(t, new(CarrierDBTestSuite))
}

func (s *CarrierDBTestSuite) TestSave() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	carrierDB := &entities.Carrier{
		ID:              uint(time.Now().Unix()),
		Name:            "Teste Carrier",
		Service:         "Entrega Expressa",
		Deadline:        3,
		Price:           120.50,
		QuoteResponseID: uint(time.Now().Unix()),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Salva o Carrier no banco de dados
	err := s.carrierDB.Save(ctx, carrierDB)
	s.Nil(err)

	// Verifica se o registro foi salvo corretamente
	var count int64
	err = s.db.Model(&entities.Carrier{}).Where("id = ?", carrierDB.ID).Count(&count).Error
	s.Nil(err)
	s.Equal(int64(1), count)

}
