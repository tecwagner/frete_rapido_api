package createquote

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/frete_rapido_api/internal/useCase/mocks"
)

func TestCreateQuoteUseCase_Execute_Success(t *testing.T) {
	ctx := context.Background()
	mockGateway := new(mocks.MockCarrierGateway)

	mockQuoteFetcher := func(input CreateQuoteInputDTO) (FreightFastOutputDTO, error) {
		return FreightFastOutputDTO{
			Dispatchers: []struct {
				Offers []Offer `json:"offers"`
			}{
				{
					Offers: []Offer{
						{
							Carrier: struct {
								Name string `json:"name"`
							}{Name: "Carrier1"},
							Service: "Standard",
							DeliveryTime: struct {
								Days int `json:"days"`
							}{Days: 5},
							FinalPrice: 100.0,
						},
					},
				},
			},
		}, nil
	}

	useCase := NewCreateQuoteUseCase(mockGateway, mockQuoteFetcher)

	input := CreateQuoteInputDTO{
		Shipper:   Shipper{RegisteredNumber: "25438296000158", Token: "1d52a9b6b78cf07b08586152459a5c90", PlatformCode: "5AKVkHqCn"},
		Recipient: Recipient{Type: 1, Country: "BR", Zipcode: 12345},
		Dispatchers: []Dispatcher{
			{
				RegisteredNumber: "987654321",
				Zipcode:          54321,
				TotalPrice:       100.0,
				Volumes: []Volume{
					{Amount: 1, Category: "Box", Height: 10, Width: 10, Length: 10, UnitaryPrice: 50, UnitaryWeight: 5},
				},
			},
		},
		SimulationType: []int{1},
	}

	// Mock para a chamada Save que retorna nil
	mockGateway.On("Save", mock.Anything, mock.AnythingOfType("[]entities.Carrier"), mock.AnythingOfType("uint")).Return(nil).Once()

	output, err := useCase.Execute(ctx, input)

	// Verifique as asserções
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.Carriers, 1)
	assert.Equal(t, "Carrier1", output.Carriers[0].Name)

	// Verifique se o método Save foi chamado com os parâmetros corretos
	mockGateway.AssertExpectations(t)
}

// Teste de erro ao salvar
func TestCreateQuoteUseCase_SaveError(t *testing.T) {

	ctx := context.Background()
	mockGateway := new(mocks.MockCarrierGateway)

	// Simula o comportamento da função de cotação
	mockQuoteFetcher := func(input CreateQuoteInputDTO) (FreightFastOutputDTO, error) {
		return FreightFastOutputDTO{
			Dispatchers: []struct {
				Offers []Offer `json:"offers"`
			}{
				{
					Offers: []Offer{
						{
							Carrier: struct {
								Name string `json:"name"`
							}{Name: "Carrier1"},
							Service: "Standard",
							DeliveryTime: struct {
								Days int `json:"days"`
							}{Days: 5},
							FinalPrice: 100.0,
						},
					},
				},
			},
		}, nil
	}

	// Crie o caso de uso com o mock do gateway e o mock da função de cotação
	useCase := NewCreateQuoteUseCase(mockGateway, mockQuoteFetcher)

	input := CreateQuoteInputDTO{
		Shipper:   Shipper{RegisteredNumber: "123456789", Token: "token", PlatformCode: "platform"},
		Recipient: Recipient{Type: 1, Country: "BR", Zipcode: 12345},
		Dispatchers: []Dispatcher{
			{
				RegisteredNumber: "987654321",
				Zipcode:          54321,
				TotalPrice:       100.0,
				Volumes: []Volume{
					{Amount: 1, Category: "Box", Height: 10, Width: 10, Length: 10, UnitaryPrice: 50, UnitaryWeight: 5},
				},
			},
		},
		SimulationType: []int{1},
	}

	mockGateway.On("Save", mock.Anything, mock.AnythingOfType("[]entities.Carrier"), mock.AnythingOfType("uint")).Return(errors.New("database error")).Once()

	output, err := useCase.Execute(ctx, input)

	// Verifique as asserções
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.Equal(t, "failed to save carriers: database error", err.Error())

	// Verifique se o método Save foi chamado
	mockGateway.AssertExpectations(t)
}
