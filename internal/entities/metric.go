package entities

type CarrierMetrics struct {
	CarrierName  string  `json:"carrier_name"`
	Count        int     `json:"count"`
	TotalFreight float64 `json:"total_final_price"`
	AvgFreight   float64 `json:"avg_final_price"`
}

type MetricsResponse struct {
	CarrierMetrics       []CarrierMetrics `json:"carrier_metrics"`
	CheapestFreight      float64          `json:"cheapest_freight"`
	MostExpensiveFreight float64          `json:"most_expensive_freight"`
}

func NewMetricsResponse(carrierMetrics []CarrierMetrics, cheapestFreight, mostExpensiveFreight float64) *MetricsResponse {
	metricsResponse := &MetricsResponse{
		CarrierMetrics:       carrierMetrics,
		CheapestFreight:      cheapestFreight,
		MostExpensiveFreight: mostExpensiveFreight,
	}

	return metricsResponse
}

func (c *CarrierMetrics) CalculateAverageFreight() {
	if c.Count > 0 {
		c.AvgFreight = c.TotalFreight / float64(c.Count)
	}
}

func (m *MetricsResponse) CalculateFreightExtremes() {
	if len(m.CarrierMetrics) == 0 {
		m.CheapestFreight = 0
		m.MostExpensiveFreight = 0
		return
	}

	m.CheapestFreight = m.CarrierMetrics[0].AvgFreight
	m.MostExpensiveFreight = m.CarrierMetrics[0].AvgFreight

	for _, metric := range m.CarrierMetrics {
		if metric.AvgFreight < m.CheapestFreight {
			m.CheapestFreight = metric.AvgFreight
		}
		if metric.AvgFreight > m.MostExpensiveFreight {
			m.MostExpensiveFreight = metric.AvgFreight
		}
	}

}
