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

	mockQuoteFetcher := func(ctx context.Context, input CreateQuoteInputDTO) (FreightFastOutputDTO, error) {
		return FreightFastOutputDTO{
			Dispatchers: []Dispatcher{
				{
					Offers: []Offer{
						{
							Carrier: Carrier{
								Name: "Carrier1",
							},
							Service: "Standard",
							DeliveryTime: DeliveryTime{
								Days: 5,
							},
							FinalPrice: 100.0,
						},
					},
				},
			},
		}, nil
	}

	useCase := NewCreateQuoteUseCase(mockGateway, mockQuoteFetcher)

	input := CreateQuoteInputDTO{
		Recipient: Recipient{
			Address: Address{
				Zipcode: "12345",
			},
		},
		Volumes: []Volume{
			{Category: 1, Amount: 1, UnitaryWeight: 5, Price: 50, SKU: "SKU1", Height: 10, Width: 10, Length: 10},
		},
	}

	mockGateway.On("Save", mock.Anything, mock.AnythingOfType("[]entities.Carrier"), mock.AnythingOfType("uint")).Return(nil).Once()

	output, err := useCase.Execute(ctx, input)

	assert.NoError(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output.Carriers, 1)
	assert.Equal(t, "Carrier1", output.Carriers[0].Name)
	assert.False(t, output.NoCarriers)

	mockGateway.AssertExpectations(t)
}

func TestCreateQuoteUseCase_NoCarriers(t *testing.T) {
	ctx := context.Background()
	mockGateway := new(mocks.MockCarrierGateway)

	mockQuoteFetcher := func(ctx context.Context, input CreateQuoteInputDTO) (FreightFastOutputDTO, error) {
		return FreightFastOutputDTO{}, errors.New("there are no carriers available for this zipcode")
	}

	useCase := NewCreateQuoteUseCase(mockGateway, mockQuoteFetcher)

	input := CreateQuoteInputDTO{
		Recipient: Recipient{
			Address: Address{
				Zipcode: "12345",
			},
		},
		Volumes: []Volume{
			{Category: 1, Amount: 1, UnitaryWeight: 5, Price: 50, SKU: "SKU1", Height: 10, Width: 10, Length: 10},
		},
	}

	output, err := useCase.Execute(ctx, input)

	assert.Error(t, err)
	assert.EqualError(t, err, err.Error())
	assert.Nil(t, output)
}

func TestCreateQuoteUseCase_SaveError(t *testing.T) {
	ctx := context.Background()
	mockGateway := new(mocks.MockCarrierGateway)

	mockQuoteFetcher := func(ctx context.Context, input CreateQuoteInputDTO) (FreightFastOutputDTO, error) {
		return FreightFastOutputDTO{
			Dispatchers: []Dispatcher{
				{
					Offers: []Offer{
						{
							Carrier: Carrier{
								Name: "Carrier1",
							},
							Service: "Standard",
							DeliveryTime: DeliveryTime{
								Days: 5,
							},
							FinalPrice: 100.0,
						},
					},
				},
			},
		}, nil
	}

	useCase := NewCreateQuoteUseCase(mockGateway, mockQuoteFetcher)

	input := CreateQuoteInputDTO{
		Recipient: Recipient{
			Address: Address{
				Zipcode: "12345",
			},
		},
		Volumes: []Volume{
			{Category: 1, Amount: 1, UnitaryWeight: 5, Price: 50, SKU: "SKU1", Height: 10, Width: 10, Length: 10},
		},
	}

	mockGateway.On("Save", mock.Anything, mock.AnythingOfType("[]entities.Carrier"), mock.AnythingOfType("uint")).Return(errors.New("database error")).Once()

	output, err := useCase.Execute(ctx, input)

	assert.Error(t, err)
	assert.Nil(t, output)
	assert.Equal(t, "failed to save carriers: database error", err.Error())

	mockGateway.AssertExpectations(t)
}
