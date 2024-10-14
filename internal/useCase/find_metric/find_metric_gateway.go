package findmetric

import "github.com/tecwagner/frete_rapido_api/internal/gateway"

type FindMentricUseCase struct {
	metricGateway gateway.IMetricGateway
}
