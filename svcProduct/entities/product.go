package entities

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// Validate validates the PaymentRequest fields.
func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(5, 255)),
		validation.Field(&p.Price, validation.Required),
		validation.Field(&p.Description, validation.Required, validation.Length(10, 1000)),
	)
}
