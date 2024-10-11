package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCarrier(t *testing.T) {

	t.Run("casos de sucesso", func(t *testing.T) {
		carrier, err := NewCarrier("teste", "teste", 3, 10.00, 1)

		assert.Nil(t, err)
		assert.NotNil(t, carrier)
		assert.Equal(t, "teste", carrier.Name)
		assert.Equal(t, "teste", carrier.Service)
		assert.Equal(t, 3, carrier.Deadline)
		assert.Equal(t, 10.00, carrier.Price)
	})

	t.Run("casos de falha", func(t *testing.T) {

		test := []struct {
			testName          string
			name              string
			service           string
			deadline          int
			price             float64
			quote_response_id int
			expectedErr       string
		}{
			{"should return error when name is empty", "", "teste", 3, 10.00, 1, "name is required"},
			{"should return error when service is empty", "teste", "", 3, 10.00, 1, "service is required"},
			{"should return error when deadline is empty", "teste", "teste", 0, 10.00, 1, "deadline must be greater than 0"},
			{"should return error when price is empty", "teste", "teste", 3, -1, 1, "price must be greater than 0"},
		}

		for _, tt := range test {
			t.Run(tt.name, func(t *testing.T) {
				carrier, err := NewCarrier(tt.name, tt.service, tt.deadline, tt.price, uint(tt.quote_response_id))
				assert.Nil(t, carrier)
				assert.Equal(t, tt.expectedErr, err.Error())
				assert.Empty(t, carrier)
			})
		}

	})

}
