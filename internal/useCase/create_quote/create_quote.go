package createquote

import (
	"context"
	"errors"
	"time"

	"github.com/tecwagner/frete_rapido_api/internal/entities"
	"github.com/tecwagner/frete_rapido_api/internal/gateway"
)

func NewCreateQuoteUseCase(quoteGateway gateway.ICarrierGateway, quoteFetcher QuoteFetcherFunc) *CreateQuoteUseCase {
	return &CreateQuoteUseCase{
		quoteGateway: quoteGateway,
		quoteFetcher: quoteFetcher,
	}
}

func (uc *CreateQuoteUseCase) Execute(ctx context.Context, input CreateQuoteInputDTO) (*CreateQuoteOutputDTO, error) {

	freightResponse, err := uc.quoteFetcher(input)
	if err != nil {
		return nil, errors.New("failed to get quote from Frete Rápido API: " + err.Error())
	}

	quoteResponse := ProcessFreightFastResponse(freightResponse)

	for _, carrierDTO := range quoteResponse.Carriers {
		carrierEntity := &entities.Carrier{
			Name:      carrierDTO.Name,
			Service:   carrierDTO.Service,
			Deadline:  carrierDTO.Deadline,
			Price:     carrierDTO.Price,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = uc.quoteGateway.Save(ctx, carrierEntity)
		if err != nil {
			return nil, errors.New("failed to save quote: " + err.Error())
		}
	}

	return quoteResponse, nil
}

func ProcessFreightFastResponse(input FreightFastOutputDTO) *CreateQuoteOutputDTO {
	var carriers []Carrier

	for _, dispatcher := range input.Dispatchers {
		for _, offer := range dispatcher.Offers {
			carrier := Carrier{
				Name:     offer.Carrier.Name,
				Service:  offer.Service,
				Deadline: offer.DeliveryTime.Days,
				Price:    offer.FinalPrice,
			}
			carriers = append(carriers, carrier)
		}
	}

	return &CreateQuoteOutputDTO{
		Carriers: carriers,
	}
}
