// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package sqlc

import (
	"time"

	zero "gopkg.in/guregu/null.v4/zero"
)

type Account struct {
	ID           string    `json:"id"`
	FullName     string    `json:"full_name"`
	BirthdayDate time.Time `json:"birthday_date"`
	AddedDate    time.Time `json:"added_date"`
	Role         string    `json:"role"`
	PhoneNumber  string    `json:"phone_number"`
	Password     string    `json:"password"`
	CreatedAt    zero.Time `json:"created_at"`
	UpdatedAt    zero.Time `json:"updated_at"`
	DeletedAt    zero.Time `json:"deleted_at"`
}
