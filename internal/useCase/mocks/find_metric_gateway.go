package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/frete_rapido_api/internal/entities"
)

type MockMetricGateway struct {
	mock.Mock
}

func (m *MockMetricGateway) Find(ctx context.Context, lastQuotes *int) (*entities.MetricsResponse, error) {
	args := m.Called(ctx, lastQuotes)
	return args.Get(0).(*entities.MetricsResponse), args.Error(1)
}
