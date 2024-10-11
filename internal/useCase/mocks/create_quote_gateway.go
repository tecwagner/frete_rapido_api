package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/frete_rapido_api/internal/entities"
)

type MockCarrierGateway struct {
	mock.Mock
}

func (m *MockCarrierGateway) Save(ctx context.Context, carriers []entities.Carrier, quoteResponseID uint) error {
	args := m.Called(ctx, carriers, quoteResponseID)
	return args.Error(0)
}
