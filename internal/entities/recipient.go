package entities

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Recipient struct {
	Type    int    `json:"type"`
	Country string `json:"country"`
	Zipcode int    `json:"zipcode"`
}

func NewRecipient(typeCode int, country string, zipcode int) (*Recipient, error) {
	recipient := &Recipient{
		Type:    typeCode,
		Country: country,
		Zipcode: zipcode,
	}

	if err := recipient.Validate(); err != nil {
		return nil, err
	}

	return recipient, nil
}

func (r *Recipient) Validate() error {

	if err := r.validateType(); err != nil {
		return err
	}

	if err := r.validateCoutry(); err != nil {
		return err
	}

	if err := r.validateZipcode(); err != nil {
		return err
	}

	return nil
}

func (r *Recipient) validateType() error {
	if r.Type < 0 || r.Type > 1 {
		return errors.New("type must be either 0 or 1")
	}
	return nil
}

func (r *Recipient) validateCoutry() error {

	if r.Country == "" {
		return errors.New("country is required")
	} else if len(r.Country) != 3 {
		return errors.New("country must be exactly 3 characters long")
	}
	return nil
}

func (r *Recipient) validateZipcode() error {
	cepString := strconv.Itoa(r.Zipcode)

	if r.Country == "BRA" {
		regexCEP := regexp.MustCompile(`^\d{8}$`)
		if !regexCEP.MatchString(cepString) {
			return fmt.Errorf("invalid Brazilian zipcode: %s", cepString)
		}
	}

	return nil
}
