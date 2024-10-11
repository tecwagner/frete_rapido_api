package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRecipient(t *testing.T) {

	t.Run("casos de sucesso", func(t *testing.T) {
		testCases := []struct {
			name        string
			types       int
			country     string
			zipcode     int
			expectedErr string
		}{
			{
				name:        "type pessoa fisica",
				types:       0,
				country:     "BRA",
				zipcode:     12345678,
				expectedErr: "",
			},
			{
				name:        "type pessoa juridica",
				types:       1,
				country:     "BRA",
				zipcode:     12345678,
				expectedErr: "",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				recipient, err := NewRecipient(tc.types, tc.country, tc.zipcode)
				assert.Nil(t, err)
				assert.NotNil(t, recipient)
				assert.Equal(t, tc.types, recipient.Type)
				assert.Equal(t, tc.country, recipient.Country)
				assert.Equal(t, tc.zipcode, recipient.Zipcode)
			})
		}
	})

	t.Run("casos de falha", func(t *testing.T) {
		testCases := []struct {
			name        string
			types       int
			country     string
			zipcode     int
			expectedErr string
		}{
			{
				name:        "type inválido",
				types:       -1,
				country:     "BRA",
				zipcode:     12345678,
				expectedErr: "type must be either 0 or 1",
			},
			{
				name:        "country inválido",
				types:       1,
				country:     "",
				zipcode:     12345678,
				expectedErr: "country is required",
			},
			{
				name:        "zipcode inválido",
				types:       1,
				country:     "BRA",
				zipcode:     -1,
				expectedErr: "invalid Brazilian zipcode: -1",
			},
			{
				name:        "coutry 3 caracteres",
				types:       1,
				country:     "BRAR",
				zipcode:     123456789,
				expectedErr: "country must be exactly 3 characters long",
			},
			{
				name:        "coutry 4 caracteres",
				types:       1,
				country:     "BRAS",
				zipcode:     123456789,
				expectedErr: "country must be exactly 3 characters long",
			},
			{
				name:        "coutry 2 caracteres",
				types:       1,
				country:     "BR",
				zipcode:     123456789,
				expectedErr: "country must be exactly 3 characters long",
			},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				recipient, err := NewRecipient(tc.types, tc.country, tc.zipcode)
				assert.NotNil(t, err)
				assert.EqualError(t, err, tc.expectedErr)
				assert.Empty(t, recipient)
			})
		}
	})
}
