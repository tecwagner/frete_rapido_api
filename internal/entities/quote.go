package entities

import "errors"

type QuoteRequest struct {
	Shipper        Shipper      `json:"shipper"`
	Recipient      Recipient    `json:"recipient"`
	Dispatchers    []Dispatcher `json:"dispatchers"`
	SimulationType []int        `json:"simulation_type"`
}

func NewQuoteRequest(shipper Shipper, recipient Recipient, dispatchers []Dispatcher, simulationType []int) (*QuoteRequest, error) {
	quote := &QuoteRequest{
		Shipper:        shipper,
		Recipient:      recipient,
		Dispatchers:    dispatchers,
		SimulationType: simulationType,
	}

	if err := quote.Validate(); err != nil {
		return nil, err
	}

	return quote, nil
}

func (sq *QuoteRequest) Validate() error {

	if err := sq.Shipper.Validate(); err != nil {
		return err
	}

	if err := sq.Recipient.Validate(); err != nil {
		return err
	}

	if len(sq.Dispatchers) == 0 {
		return errors.New("at least one dispatcher is required")
	}

	for _, dispatcher := range sq.Dispatchers {
		if err := dispatcher.Validate(); err != nil {
			return err
		}
	}

	if err := sq.validateSimulationType(); err != nil {
		return err
	}

	return nil
}

func (sq *QuoteRequest) validateSimulationType() error {
	if len(sq.SimulationType) == 0 {
		return errors.New("at least one simulation type is required")
	}

	for _, simType := range sq.SimulationType {
		if simType < 0 || simType > 1 {
			return errors.New("simulation type must be either 0 or 1")
		}
	}

	return nil
}
