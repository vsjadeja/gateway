package entities

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate validates the PaymentRequest fields.
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required, validation.Length(5, 50)),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 16)),
	)
}
