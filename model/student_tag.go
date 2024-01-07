package model
import "github.com/google/uuid"
type CreateStudentTag struct {
	Title       string `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	Description string `json:"description" db:"description" lenMin:"0" lenMax:"1024" regex:"login"`
}
type StudentTag struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	Description string    `json:"description" db:"description" lenMin:"0" lenMax:"1024" regex:"login"`
}
