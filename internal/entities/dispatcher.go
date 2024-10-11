package entities

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          int      `json:"zipcode"`
	TotalPrice       float64  `json:"total_price"`
	Volumes          []Volume `json:"volumes"`
}

func NewDispatcher(registeredNumber string, zipcode int, totalPrice float64, volumes []Volume) (*Dispatcher, error) {
	dispacher := &Dispatcher{
		RegisteredNumber: registeredNumber,
		Zipcode:          zipcode,
		TotalPrice:       totalPrice,
		Volumes:          volumes,
	}

	err := dispacher.Validate()
	if err != nil {
		return nil, err
	}

	return dispacher, nil
}

func (d *Dispatcher) Validate() error {

	if d.RegisteredNumber == "" {
		return errors.New("registered number is required")
	} else if len(d.RegisteredNumber) < 14 {
		return errors.New("registered number must be exactly 14 characters long")
	}

	if err := d.validateZipcode(); err != nil {
		return err
	}

	if d.TotalPrice < 0.0 {
		return errors.New("total price must be greater than zero")
	}

	if len(d.Volumes) == 0 {
		return errors.New("at least one volume is required")
	}

	for _, volume := range d.Volumes {
		if err := volume.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (d *Dispatcher) validateZipcode() error {
	cepString := strconv.Itoa(d.Zipcode)

	regexCEP := regexp.MustCompile(`^\d{8}$`)
	if !regexCEP.MatchString(cepString) {
		return fmt.Errorf("invalid Brazilian zipcode: %s", cepString)
	}

	return nil
}
