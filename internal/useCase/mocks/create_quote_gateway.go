package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/frete_rapido_api/internal/entities"
)

type MockCarrierGateway struct {
	mock.Mock
}

func (m *MockCarrierGateway) Save(ctx context.Context, carrier *entities.Carrier) error {
	args := m.Called(carrier)
	return args.Error(0)
}
