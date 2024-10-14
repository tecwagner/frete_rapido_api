package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetricsResponse_Success(t *testing.T) {
	carrierMetrics := []CarrierMetrics{
		{
			CarrierName:  "Carrier A",
			Count:        10,
			TotalFreight: 1000,
		},
		{
			CarrierName:  "Carrier B",
			Count:        5,
			TotalFreight: 750,
		},
	}
	metricsResponse := NewMetricsResponse(carrierMetrics, 0, 0)

	for i := range metricsResponse.CarrierMetrics {
		metricsResponse.CarrierMetrics[i].CalculateAverageFreight()
	}

	metricsResponse.CalculateFreightExtremes()

	assert.Equal(t, 100.0, metricsResponse.CarrierMetrics[0].AvgFreight)
	assert.Equal(t, 150.0, metricsResponse.CarrierMetrics[1].AvgFreight)
	assert.Equal(t, 100.0, metricsResponse.CheapestFreight)
	assert.Equal(t, 150.0, metricsResponse.MostExpensiveFreight)
}

func TestMetricsResponse_EmptyCarrierMetrics(t *testing.T) {
	carrierMetrics := []CarrierMetrics{}

	metricsResponse := NewMetricsResponse(carrierMetrics, 0, 0)

	metricsResponse.CalculateFreightExtremes()

	assert.Equal(t, float64(0), metricsResponse.CheapestFreight)
	assert.Equal(t, float64(0), metricsResponse.MostExpensiveFreight)
}

func TestMetricsResponse_NegativeFreight(t *testing.T) {
	carrierMetrics := []CarrierMetrics{
		{
			CarrierName:  "Carrier A",
			Count:        5,
			TotalFreight: -500,
		},
	}

	metricsResponse := NewMetricsResponse(carrierMetrics, 0, 0)

	for i := range metricsResponse.CarrierMetrics {
		metricsResponse.CarrierMetrics[i].CalculateAverageFreight()
	}

	// Cálculo do frete mais barato e mais caro
	metricsResponse.CalculateFreightExtremes()

	// Verificações
	assert.Equal(t, -100.0, metricsResponse.CheapestFreight)
	assert.Equal(t, -100.0, metricsResponse.MostExpensiveFreight)
}
