package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuoteRequest(t *testing.T) {
	t.Run("casos de sucesso", func(t *testing.T) {
		testCases := []struct {
			name           string
			shipper        Shipper
			recipient      Recipient
			dispatchers    []Dispatcher
			simulationType []int
			expectedErr    string
		}{
			{
				name: "create shipping quote with valid data",
				shipper: Shipper{
					RegisteredNumber: "21321313131356",
					Token:            "21321313131356654789632541236547",
					PlatformCode:     "XYZ",
				},
				recipient: Recipient{
					Type:    1,
					Country: "BRA",
					Zipcode: 12345678,
				},
				dispatchers: []Dispatcher{
					{
						RegisteredNumber: "21321313131356",
						Zipcode:          12345678,
						TotalPrice:       0.0,
						Volumes: []Volume{
							{
								Amount:        1,
								Category:      "7",
								Sku:           "abc-test-12",
								Height:        1.0,
								Width:         1.0,
								Length:        1.0,
								UnitaryPrice:  55.0,
								UnitaryWeight: 6.0,
							},
						},
					},
				},
				simulationType: []int{0},
				expectedErr:    "",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				quote, err := NewQuoteRequest(tc.shipper, tc.recipient, tc.dispatchers, tc.simulationType)
				assert.Nil(t, err)
				assert.NotNil(t, quote)
				assert.Equal(t, tc.shipper, quote.Shipper)
				assert.Equal(t, tc.recipient, quote.Recipient)
				assert.Equal(t, tc.dispatchers, quote.Dispatchers)
				assert.Equal(t, tc.simulationType, quote.SimulationType)
			})
		}
	})

	t.Run("casos de falha", func(t *testing.T) {
		testCases := []struct {
			name           string
			shipper        Shipper
			recipient      Recipient
			dispatchers    []Dispatcher
			simulationType []int
			expectedErr    string
		}{
			{
				name: "sem dispatchers",
				shipper: Shipper{
					RegisteredNumber: "21321313131356",
					Token:            "21321313131356654789632541236547",
					PlatformCode:     "XYZ",
				},
				recipient: Recipient{
					Type:    1,
					Country: "BRA",
					Zipcode: 12345678,
				},
				dispatchers:    []Dispatcher{},
				simulationType: []int{0},
				expectedErr:    "at least one dispatcher is required",
			},
			{
				name: "sem simulation type",
				shipper: Shipper{
					RegisteredNumber: "21321313131356",
					Token:            "21321313131356654789632541236547",
					PlatformCode:     "XYZ",
				},
				recipient: Recipient{
					Type:    1,
					Country: "BRA",
					Zipcode: 12345678,
				},
				dispatchers: []Dispatcher{
					{
						RegisteredNumber: "21321313131356",
						Zipcode:          12345678,
						TotalPrice:       0.0,
						Volumes: []Volume{
							{
								Amount:        1,
								Category:      "7",
								Sku:           "abc-test-12",
								Height:        1.0,
								Width:         1.0,
								Length:        1.0,
								UnitaryPrice:  55.0,
								UnitaryWeight: 6.0,
							},
						},
					},
				},
				simulationType: []int{},
				expectedErr:    "at least one simulation type is required",
			},
			{
				name: "simulation type inválido",
				shipper: Shipper{
					RegisteredNumber: "21321313131356",
					Token:            "21321313131356654789632541236547",
					PlatformCode:     "XYZ",
				},
				recipient: Recipient{
					Type:    1,
					Country: "BRA",
					Zipcode: 12345678,
				},
				dispatchers: []Dispatcher{
					{
						RegisteredNumber: "21321313131356",
						Zipcode:          12345678,
						TotalPrice:       0.0,
						Volumes: []Volume{
							{
								Amount:        1,
								Category:      "7",
								Sku:           "abc-test-12",
								Height:        1.0,
								Width:         1.0,
								Length:        1.0,
								UnitaryPrice:  55.0,
								UnitaryWeight: 6.0,
							},
						},
					},
				},
				simulationType: []int{2},
				expectedErr:    "simulation type must be either 0 or 1",
			},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				quote, err := NewQuoteRequest(tc.shipper, tc.recipient, tc.dispatchers, tc.simulationType)
				assert.NotNil(t, err)
				assert.EqualError(t, err, tc.expectedErr)
				assert.Nil(t, quote) // O quote deve ser nil em caso de erro
			})
		}
	})
}
