package gateway

import (
	"context"

	"github.com/tecwagner/frete_rapido_api/internal/entities"
)

type IMetricGateway interface {
	Find(ctx context.Context, lastQuotes *int) (*entities.MetricsResponse, error)
}
