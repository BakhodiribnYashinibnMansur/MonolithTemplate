package model

import (
	"github.com/google/uuid"
	"time"
)

type Manager struct {
	ID           uuid.UUID `json:"id"  db:"id"`
	FullName     string    `json:"fullName" db:"full_name"`
	BirthdayDate string    `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string    `json:"phoneNumber" db:"phone_number"`
	RoleID       uuid.UUID `json:"-" db:"role_id"`
	Role         RoleData  `json:"role" db:"-"`
	PhotoName    string    `json:"-" db:"photo"`
	Photo        File      `json:"photo" db:"-"`
	AddedDate    string    `json:"addedDate" db:"added_date"`
}
type CreateManager struct {
	ID           uuid.UUID `json:"-"  db:"id"`
	FullName     string    `json:"fullName" db:"full_name" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string    `json:"phoneNumber" db:"phone_number" required:"true" lenMin:"0" lenMax:"16" regex:"phone"`
	Password     string    `json:"password" db:"password"  lenMin:"8" lenMax:"64"`
	RoleID       string    `json:"-" db:"role_id"`
	RoleTitle    string    `json:"-" db:"-"`
	Photo        string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	AddedDate    time.Time `json:"addedDate" db:"added_date" required:"true"`
}
type UpdateManager struct {
	ID           uuid.UUID `json:"-"  db:"id"`
	FullName     string    `json:"fullName" db:"full_name" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string    `json:"phoneNumber" db:"phone_number" required:"true" lenMin:"0" lenMax:"16" regex:"phone"`
	RoleID       string    `json:"-" db:"role_id"`
	RoleTitle    string    `json:"-" db:"-"`
	Photo        string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	AddedDate    time.Time `json:"addedDate" db:"added_date" required:"true"`
}
