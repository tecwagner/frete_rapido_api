package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDispatcher(t *testing.T) {

	t.Run("casos de sucesso", func(t *testing.T) {
		dispatcher, err := NewDispatcher("21321313131356", 12345678, 0.0, []Volume{
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
			{
				Amount:        1,
				Category:      "7",
				Sku:           "abc-test-34",
				Height:        1.0,
				Width:         1.0,
				Length:        1.0,
				UnitaryPrice:  56.0,
				UnitaryWeight: 6.0,
			},
		})
		assert.Nil(t, err)
		assert.NotNil(t, dispatcher)
		assert.Equal(t, "21321313131356", dispatcher.RegisteredNumber)
		assert.Equal(t, 12345678, dispatcher.Zipcode)
		assert.Equal(t, 0.0, dispatcher.TotalPrice)
		assert.Equal(t, 2, len(dispatcher.Volumes))

	})

	t.Run("casos de falhas", func(t *testing.T) {

		testCases := []struct {
			name             string
			registeredNumber string
			zipcode          int
			totalPrice       float64
			volumes          []Volume
			expectedErr      string
		}{
			{
				name:             "registered Number inválido",
				registeredNumber: "",
				zipcode:          12345678,
				totalPrice:       0.0,
				volumes: []Volume{
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
				expectedErr: "registered number is required",
			},
			{
				name:             "Registered number with less than 14 characters",
				registeredNumber: "123",
				zipcode:          12345678,
				totalPrice:       0.0,
				volumes: []Volume{
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
				expectedErr: "registered number must be exactly 14 characters long",
			},
			{
				name:             "zipcode inválido",
				registeredNumber: "21321313131356",
				zipcode:          0,
				totalPrice:       0.0,
				volumes: []Volume{
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
				expectedErr: "invalid Brazilian zipcode: 0",
			},
			{
				name:             "totalPrice inválido",
				registeredNumber: "21321313131356",
				zipcode:          12345678,
				totalPrice:       -1.0,
				volumes: []Volume{
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
				expectedErr: "total price must be greater than zero",
			},
			{
				name:             "volumes inválido",
				registeredNumber: "21321313131356",
				zipcode:          12345678,
				totalPrice:       0.0,
				volumes:          []Volume{},
				expectedErr:      "at least one volume is required",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				dispatcher, err := NewDispatcher(tc.registeredNumber, tc.zipcode, tc.totalPrice, tc.volumes)
				assert.NotNil(t, err)
				assert.EqualError(t, err, tc.expectedErr)
				assert.Empty(t, dispatcher)
			})
		}

	})
}
