package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVolume(t *testing.T) {

	t.Run("casos de sucessos", func(t *testing.T) {
		volume, err := NewVolume(1, "7", "abc-test", 0.2, 0.2, 0.2, 55.0, 6.0)

		assert.Nil(t, err)
		assert.NotNil(t, volume)
		assert.Equal(t, 1, volume.Amount)
		assert.Equal(t, "7", volume.Category)
		assert.Equal(t, "abc-test", volume.Sku)
		assert.Equal(t, 0.2, volume.Width)
		assert.Equal(t, 0.2, volume.Height)
		assert.Equal(t, 0.2, volume.Length)
		assert.Equal(t, 55.0, volume.UnitaryPrice)
		assert.Equal(t, 6.0, volume.UnitaryWeight)
	})

	t.Run("casos de falhas", func(t *testing.T) {

		testCases := []struct {
			name          string
			amount        int
			category      string
			sku           string
			height        float64
			width         float64
			length        float64
			unitaryPrice  float64
			unitaryWeight float64
			expectedErr   string
		}{
			{
				name:          "amount invalid",
				amount:        0,
				category:      "7",
				sku:           "abc-test",
				height:        0.2,
				width:         0.2,
				length:        0.2,
				unitaryPrice:  55,
				unitaryWeight: 6,
				expectedErr:   "amount must be greater than 0",
			},

			{
				name:          "category invalid",
				amount:        1,
				category:      "",
				sku:           "abc-test",
				height:        0.2,
				width:         0.2,
				length:        0.2,
				unitaryPrice:  55,
				unitaryWeight: 6,
				expectedErr:   "category is required",
			},
			{
				name:          "sku invalid",
				amount:        1,
				category:      "7",
				sku:           "",
				height:        0.2,
				width:         0.2,
				length:        0.2,
				unitaryPrice:  55,
				unitaryWeight: 6,
				expectedErr:   "sku is required",
			},
			{
				name:          "height invalid",
				amount:        1,
				category:      "7",
				sku:           "abc-test",
				height:        0,
				width:         0.2,
				length:        0.2,
				unitaryPrice:  55,
				unitaryWeight: 6,
				expectedErr:   "height must be greater than 0",
			},
			{
				name:          "width invalid",
				amount:        1,
				category:      "7",
				sku:           "abc-test",
				height:        0.2,
				width:         0,
				length:        0.2,
				unitaryPrice:  55,
				unitaryWeight: 6,
				expectedErr:   "width must be greater than 0",
			},
			{
				name:          "length invalid",
				amount:        1,
				category:      "7",
				sku:           "abc-test",
				height:        0.2,
				width:         0.2,
				length:        0,
				unitaryPrice:  55,
				unitaryWeight: 6,
				expectedErr:   "length must be greater than 0",
			},
			{
				name:          "unitaryPrice invalid",
				amount:        1,
				category:      "7",
				sku:           "abc-test",
				height:        0.2,
				width:         0.2,
				length:        0.2,
				unitaryPrice:  0,
				unitaryWeight: 6,
				expectedErr:   "unitary_price must be greater than 0",
			},
			{
				name:          "unitaryWeight invalid",
				amount:        1,
				category:      "7",
				sku:           "abc-test",
				height:        0.2,
				width:         0.2,
				length:        0.2,
				unitaryPrice:  55,
				unitaryWeight: 0,
				expectedErr:   "unitary_weight must be greater than 0",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				volume, err := NewVolume(tc.amount, tc.category, tc.sku, tc.height, tc.width, tc.length, tc.unitaryPrice, tc.unitaryWeight)
				assert.NotNil(t, err)
				assert.EqualError(t, err, tc.expectedErr)
				assert.Empty(t, volume)
			})
		}
	})
}
