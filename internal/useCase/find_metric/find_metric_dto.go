package findmetric

type CarrierMetrics struct {
	CarrierName     string  `json:"carrier_name"`
	Count           int     `json:"count"`
	TotalFinalPrice float64 `json:"total_final_price"`
	AvgFinalPrice   float64 `json:"avg_final_price"`
}

type MetricsOutputDTO struct {
	CarrierMetrics       []CarrierMetrics `json:"carrier_metrics"`
	CheapestFreight      float64          `json:"cheapest_freight"`
	MostExpensiveFreight float64          `json:"most_expensive_freight"`
}
