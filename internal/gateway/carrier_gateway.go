package gateway

import (
	"context"

	"github.com/tecwagner/frete_rapido_api/internal/entities"
)

type ICarrierGateway interface {
	Save(ctx context.Context, carrier []entities.Carrier, quoteResponseID uint) error
}
