package entities

type Product struct {
	ID          uint    `json:"id" validate:"omitempty"`
	Name        string  `json:"name" validate:"required,min=1,max=255"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description" validate:"required,min=3"`
}
