package createquote

import (
	"context"

	"github.com/tecwagner/frete_rapido_api/internal/gateway"
)

type QuoteFetcherFunc func(ctx context.Context, request CreateQuoteInputDTO) (FreightFastOutputDTO, error)

type CreateQuoteUseCase struct {
	quoteGateway gateway.ICarrierGateway
	quoteFetcher QuoteFetcherFunc
}
