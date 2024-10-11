package entities

import (
	"errors"
)

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

func NewShipper(registeredNumber string, token string, platformCode string) (*Shipper, error) {
	shipper := &Shipper{
		RegisteredNumber: registeredNumber,
		Token:            token,
		PlatformCode:     platformCode,
	}

	err := shipper.Validate()
	if err != nil {
		return nil, err
	}

	return shipper, nil
}

func (s *Shipper) Validate() error {

	if s.RegisteredNumber == "" {
		return errors.New("registered number is required")
	} else if len(s.RegisteredNumber) < 14 {
		return errors.New("registered number must be exactly 14 characters long")
	}

	if s.Token == "" {
		return errors.New("token is required")
	} else if len(s.Token) < 32 {
		return errors.New("token must be exactly 32 characters long")
	}

	if s.PlatformCode == "" {
		return errors.New("platform code is required")
	}

	return nil
}
