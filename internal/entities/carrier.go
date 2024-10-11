package entities

import (
	"errors"
	"time"
)

type Carrier struct {
	ID              uint      `json:"id" gorm:"primaryKey autoIncrementIncrement"`
	Name            string    `json:"name" gorm:"size:255;not null"`
	Service         string    `json:"service" gorm:"size:255;not null"`
	Deadline        int       `json:"deadline" gorm:"not null"`
	Price           float64   `json:"price" gorm:"not null"`
	QuoteResponseID uint      `json:"quote_response_id" gorm:"not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func NewCarrier(name string, service string, deadline int, price float64) (*Carrier, error) {
	carrier := &Carrier{
		ID:              uint(time.Now().Unix()),
		Name:            name,
		Service:         service,
		Deadline:        deadline,
		Price:           price,
		QuoteResponseID: uint(time.Now().Unix()),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := carrier.Validate(); err != nil {
		return nil, err
	}

	return carrier, nil
}

func (c *Carrier) Validate() error {

	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Service == "" {
		return errors.New("service is required")
	}

	if c.Deadline <= 0 {
		return errors.New("deadline must be greater than 0")
	}

	if c.Price < 0 {
		return errors.New("price must be greater than 0")
	}

	return nil
}
