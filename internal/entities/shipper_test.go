package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShipper(t *testing.T) {
	t.Run("casos de sucesso", func(t *testing.T) {
		shipper, err := NewShipper("21321313131356", "21321313131356654789632541236547", "56456454564")
		assert.Nil(t, err)
		assert.NotNil(t, shipper)
		assert.Equal(t, "21321313131356", shipper.RegisteredNumber)
		assert.Equal(t, "21321313131356654789632541236547", shipper.Token)
		assert.Equal(t, "56456454564", shipper.PlatformCode)
	})

	t.Run("casos de falha", func(t *testing.T) {
		testCases := []struct {
			name             string
			registeredNumber string
			token            string
			platformCode     string
			expectedErr      string
		}{
			{
				name:             "registeredNumber vazio",
				registeredNumber: "",
				token:            "21321313131356654789632541236547",
				platformCode:     "56456454564",
				expectedErr:      "registered number is required",
			},
			{
				name:             "token vazio",
				registeredNumber: "21321313131356",
				token:            "",
				platformCode:     "56456454564",
				expectedErr:      "token is required",
			},
			{
				name:             "platformCode vazio",
				registeredNumber: "21321313131356",
				token:            "21321313131356654789632541236547",
				platformCode:     "",
				expectedErr:      "platform code is required",
			},
			{
				name:             "registeredNumber menor que 14 caracteres",
				registeredNumber: "123",
				token:            "21321313131356654789632541236547",
				platformCode:     "56456454564",
				expectedErr:      "registered number must be exactly 14 characters long",
			},
			{
				name:             "token menor que 32 caracteres",
				registeredNumber: "21321313131356",
				token:            "123",
				platformCode:     "56456454564",
				expectedErr:      "token must be exactly 32 characters long",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				shipper, err := NewShipper(tc.registeredNumber, tc.token, tc.platformCode)
				assert.NotNil(t, err)
				assert.EqualError(t, err, tc.expectedErr)
				assert.Empty(t, shipper)
			})
		}
	})
}
