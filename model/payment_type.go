package model

import "github.com/google/uuid"

type CreatePaymentType struct {
	Title       string `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	Description string `json:"description" db:"description" lenMin:"0" lenMax:"1024" regex:"login"`
}
type PaymentType struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
}
type UpdatePaymentType struct {
	Id          uuid.UUID `json:"id" db:"id" `
	Title       string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64"  regex:"login"`
	Description string    `json:"description" db:"description"  lenMin:"0" lenMax:"1024"  regex:"login"`
}
