package entities

import "errors"

type Volume struct {
	Amount        int     `json:"amount"`
	Category      string  `json:"category"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
	UnitaryWeight float64 `json:"unitary_weight"`
}

func NewVolume(amount int, category string, sku string, height float64, width float64, length float64, unitaryPrice float64, unitaryWeight float64) (*Volume, error) {
	volume := &Volume{
		Amount:        amount,
		Category:      category,
		Sku:           sku,
		Height:        height,
		Width:         width,
		Length:        length,
		UnitaryPrice:  unitaryPrice,
		UnitaryWeight: unitaryWeight,
	}

	if err := volume.Validate(); err != nil {
		return nil, err
	}
	return volume, nil

}

func (v *Volume) Validate() error {

	if err := v.validateAmount(); err != nil {
		return err
	}

	if err := v.validateCategory(); err != nil {
		return err
	}

	if err := v.validateSku(); err != nil {
		return err
	}

	if err := v.validateHeight(); err != nil {
		return err
	}

	if err := v.validateWidth(); err != nil {
		return err
	}

	if err := v.validateLength(); err != nil {
		return err
	}

	if err := v.validateUnitaryPrice(); err != nil {
		return err
	}

	if err := v.validateUnitaryWeight(); err != nil {
		return err
	}

	return nil
}

func (v *Volume) validateAmount() error {

	if v.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	return nil
}

func (v *Volume) validateCategory() error {

	if v.Category == "" {
		return errors.New("category is required")
	}

	return nil
}

func (v *Volume) validateSku() error {

	if v.Sku == "" {
		return errors.New("sku is required")
	}

	return nil
}

func (v *Volume) validateHeight() error {

	if v.Height <= 0.0 {
		return errors.New("height must be greater than 0")
	}

	return nil
}

func (v *Volume) validateWidth() error {

	if v.Width <= 0.0 {
		return errors.New("width must be greater than 0")
	}

	return nil
}

func (v *Volume) validateLength() error {

	if v.Length <= 0.0 {
		return errors.New("length must be greater than 0")
	}

	return nil
}

func (v *Volume) validateUnitaryPrice() error {

	if v.UnitaryPrice <= 0.0 {
		return errors.New("unitary_price must be greater than 0")
	}

	return nil
}

func (v *Volume) validateUnitaryWeight() error {

	if v.UnitaryWeight <= 0.0 {
		return errors.New("unitary_weight must be greater than 0")
	}
	return nil
}
