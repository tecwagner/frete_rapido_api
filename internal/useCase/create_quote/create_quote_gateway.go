package createquote

import "github.com/tecwagner/frete_rapido_api/internal/gateway"

type QuoteFetcherFunc func(request CreateQuoteInputDTO) (FreightFastOutputDTO, error)

type CreateQuoteUseCase struct {
	quoteGateway gateway.ICarrierGateway
	quoteFetcher QuoteFetcherFunc
}
