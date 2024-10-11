package createquote

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/frete_rapido_api/internal/useCase/mocks"
)

// Teste do caso de uso CreateQuoteUseCase
func TestCreateQuoteUseCase_Execute(t *testing.T) {
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

	// Defina o que o mock deve retornar ao salvar
	mockGateway.On("Save", mock.AnythingOfType("*entities.Carrier")).Return(nil).Once()

	ctx := context.Background()
	output, err := useCase.Execute(ctx, input)

	// Verifique as asserções
	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.Carriers, 1)
	assert.Equal(t, "Carrier1", output.Carriers[0].Name)

	// Verifique se o método Save foi chamado
	mockGateway.AssertExpectations(t)

}

// Teste de erro ao salvar
func TestCreateQuoteUseCase_SaveError(t *testing.T) {
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

	// Simule um erro ao salvar
	mockGateway.On("Save", mock.AnythingOfType("*entities.Carrier")).Return(errors.New("database error")).Once()

	ctx := context.Background()
	output, err := useCase.Execute(ctx, input)

	// Verifique as asserções
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.Equal(t, "failed to save quote: database error", err.Error())

	// Verifique se o método Save foi chamado
	mockGateway.AssertExpectations(t)
}
