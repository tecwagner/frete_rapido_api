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

	freightResponse, err := uc.quoteFetcher(ctx, input)
	if err != nil {
		return nil, errors.New("failed to get quote from Frete Rápido API: " + err.Error())
	}

	quoteResponse := ProcessFreightFastResponse(freightResponse)
	if quoteResponse.NoCarriers {
		return nil, errors.New("here are no carriers available for this zipcode")
	}

	quoteResponseID := uint(time.Now().Unix())
	var carrierEntities []entities.Carrier

	for _, carrierDTO := range quoteResponse.Carriers {
		if carrierDTO.Name == "" || carrierDTO.Service == "" {
			return nil, errors.New("carrier DTO has empty fields")
		}
		carrierEntity := entities.Carrier{
			Name:     carrierDTO.Name,
			Service:  carrierDTO.Service,
			Deadline: carrierDTO.Deadline,
			Price:    carrierDTO.Price,
		}
		carrierEntities = append(carrierEntities, carrierEntity)
	}

	err = uc.quoteGateway.Save(ctx, carrierEntities, quoteResponseID)
	if err != nil {
		return nil, errors.New("failed to save carriers: " + err.Error())
	}

	return quoteResponse, nil
}

func ProcessFreightFastResponse(input FreightFastOutputDTO) *CreateQuoteOutputDTO {
	if len(input.Dispatchers) == 0 {
		return &CreateQuoteOutputDTO{Carriers: []Carrier{}, NoCarriers: true}
	}

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

	if len(carriers) == 0 {
		return &CreateQuoteOutputDTO{Carriers: []Carrier{}, NoCarriers: true}
	}

	return &CreateQuoteOutputDTO{
		Carriers:   carriers,
		NoCarriers: false,
	}
}
