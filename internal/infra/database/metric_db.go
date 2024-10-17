package database

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/tecwagner/frete_rapido_api/internal/entities"
	"gorm.io/gorm"
)

type MetricDB struct {
	DB *gorm.DB
}

func NewMetricDB(db *gorm.DB) *MetricDB {
	return &MetricDB{
		DB: db,
	}
}

func (m *MetricDB) Find(ctx context.Context, lastQuotes *int) (*entities.MetricsResponse, error) {
	metricsResponse := &entities.MetricsResponse{}
	quotes := []entities.Carrier{}

	query := m.DB.Order("created_at desc").Order("price asc")

	if lastQuotes != nil {
		query = query.Limit(*lastQuotes)
		err := query.Find(&quotes).Error
		if err != nil {
			return metricsResponse, fmt.Errorf("failed to find quotes: %w", err)
		}
	} else {

		var cheapestQuote entities.Carrier
		err := m.DB.Order("price asc").First(&cheapestQuote).Error
		if err != nil {
			return metricsResponse, fmt.Errorf("failed to find cheapest quote: %w", err)
		}

		var mostExpensiveQuote entities.Carrier
		err = m.DB.Order("price desc").First(&mostExpensiveQuote).Error
		if err != nil {
			return metricsResponse, fmt.Errorf("failed to find most expensive quote: %w", err)
		}

		quotes = append(quotes, cheapestQuote, mostExpensiveQuote)
	}

	if len(quotes) == 0 {
		return metricsResponse, errors.New("Find failed: no quotes found")
	}

	carrierMetricsMap := make(map[string]*entities.CarrierMetrics)

	cheapestFreight := quotes[0].Price
	mostExpensiveFreight := quotes[0].Price

	for _, quote := range quotes {
		if quote.Price < cheapestFreight {
			cheapestFreight = quote.Price
		}

		if quote.Price > mostExpensiveFreight {
			mostExpensiveFreight = quote.Price
		}

		carrierMetric, exists := carrierMetricsMap[quote.Name]
		if !exists {
			carrierMetric = &entities.CarrierMetrics{
				CarrierName:  quote.Name,
				Count:        0,
				TotalFreight: 0,
			}
			carrierMetricsMap[quote.Name] = carrierMetric
		}

		carrierMetric.Count++
		carrierMetric.TotalFreight += quote.Price
	}

	for _, metric := range carrierMetricsMap {
		metric.CalculateAverageFreight()
		metricsResponse.CarrierMetrics = append(metricsResponse.CarrierMetrics, *metric)
	}

	sort.Slice(metricsResponse.CarrierMetrics, func(i, j int) bool {
		return metricsResponse.CarrierMetrics[i].CarrierName < metricsResponse.CarrierMetrics[j].CarrierName
	})

	metricsResponse.CheapestFreight = cheapestFreight
	metricsResponse.MostExpensiveFreight = mostExpensiveFreight

	return metricsResponse, nil
}
