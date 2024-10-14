package findmetric

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tecwagner/frete_rapido_api/internal/entities"
	"github.com/tecwagner/frete_rapido_api/internal/useCase/mocks"
)

func TestMetricDB_Find_QuotesFound(t *testing.T) {
	mockMetricGateway := new(mocks.MockMetricGateway)
	mockMetricGateway.On("Find", mock.Anything, mock.AnythingOfType("*int")).Return(&entities.MetricsResponse{CarrierMetrics: []entities.CarrierMetrics{}}, nil)

	useCase := NewMetricsUseCase(mockMetricGateway)
	ctx := context.Background()

	lastQuotes := 1
	metricsResponse, err := useCase.Execute(ctx, &lastQuotes)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if metricsResponse == nil {
		t.Fatal("expected metrics response, got nil")
	}

	mockMetricGateway.AssertExpectations(t)
}

func TestMetricDB_Find_NoQuotesFound(t *testing.T) {
	mockMetricGateway := new(mocks.MockMetricGateway)
	mockMetricGateway.On("Find", mock.Anything, mock.AnythingOfType("*int")).Return((*entities.MetricsResponse)(nil), errors.New("Find failed: no quotes found"))

	useCase := NewMetricsUseCase(mockMetricGateway)
	ctx := context.Background()

	lastQuotes := 5
	_, err := useCase.Execute(ctx, &lastQuotes)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	expectedErrorMessage := "failed to find metrics: Find failed: no quotes found"
	if err.Error() != expectedErrorMessage {
		t.Fatalf("expected error message '%s', got %v", expectedErrorMessage, err)
	}

	mockMetricGateway.AssertExpectations(t)
}
