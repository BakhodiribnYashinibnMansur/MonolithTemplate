package model

import "github.com/google/uuid"

type PaymentCatalog struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
}
type CreatePaymentCatalog struct {
	Title       string `json:"title" db:"title" lenMin:"1" lenMax:"64" required:"true" regex:"login" `
	Description string `json:"description" db:"description" lenMin:"1" lenMax:"1024"  regex:"login"`
}
type UpdatePaymentCatalog struct {
	ID          uuid.UUID `json:"-" db:"id"`
	Title       string    `json:"title" db:"title" lenMin:"1" lenMax:"64" required:"true" regex:"login" `
	Description string    `json:"description" db:"description" lenMin:"1" lenMax:"1024"  regex:"login"`
}
