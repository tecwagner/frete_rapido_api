package findmetric

import (
	"context"
	"errors"
	"fmt"

	"github.com/tecwagner/frete_rapido_api/internal/gateway"
	"github.com/tecwagner/frete_rapido_api/pkg"
)

func NewMetricsUseCase(metricGateway gateway.IMetricGateway) *FindMentricUseCase {
	return &FindMentricUseCase{
		metricGateway: metricGateway,
	}
}

func (m *FindMentricUseCase) Execute(ctx context.Context, lastQuotes *int) (*MetricsOutputDTO, error) {

	if lastQuotes != nil && *lastQuotes < 0 {
		return nil, errors.New("last_quotes must be a non-negative integer")
	}

	metricsResponse, err := m.metricGateway.Find(ctx, lastQuotes)
	if err != nil {
		return nil, fmt.Errorf("failed to find metrics: %w", err)
	}

	if metricsResponse == nil {
		return &MetricsOutputDTO{}, nil
	}

	var carrierMetrics []CarrierMetrics
	for _, metric := range metricsResponse.CarrierMetrics {
		carrierMetrics = append(carrierMetrics, CarrierMetrics{
			CarrierName:     metric.CarrierName,
			Count:           metric.Count,
			TotalFinalPrice: pkg.RoundToTwoDecimals(metric.TotalFreight),
			AvgFinalPrice:   pkg.RoundToTwoDecimals(metric.AvgFreight),
		})
	}

	return &MetricsOutputDTO{
		CarrierMetrics:       carrierMetrics,
		CheapestFreight:      pkg.RoundToTwoDecimals(metricsResponse.CheapestFreight),
		MostExpensiveFreight: pkg.RoundToTwoDecimals(metricsResponse.MostExpensiveFreight),
	}, nil
}
