package model

import (
	"github.com/google/uuid"
	"time"
)

type Board struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
	List  []List    `json:"list"`
}
type CreateBoard struct {
	Title string `json:"title" db:"title" regex:"login" lenMin:"3" lenMax:"64"`
}
type UpdateBoard struct {
	ID    uuid.UUID `json:"-" db:"id"`
	Title string    `json:"title" db:"title"  regex:"login" lenMin:"3" lenMax:"64"`
}
type List struct {
	ID      uuid.UUID `json:"id" db:"id"`
	BoardID uuid.UUID `json:"-" db:"board_id"`
	Board   BoardData `json:"board" db:"-"`
	Title   string    `json:"title" db:"title"  regex:"login" lenMin:"3" lenMax:"64"`
	Lid     []Lid     `json:"lid"`
}
type CreateList struct {
	BoardID uuid.UUID `json:"boardID" db:"board_id"`
	Title   string    `json:"title" db:"title"  regex:"login" lenMin:"3" lenMax:"64"`
}
type UpdateList struct {
	ID      uuid.UUID `json:"-" db:"id"`
	BoardID uuid.UUID `json:"boardID" db:"board_id"`
	Title   string    `json:"title" db:"title"  regex:"login" lenMin:"3" lenMax:"64"`
}
type Lid struct {
	ID             uuid.UUID    `json:"id" db:"id"`
	ListID         uuid.UUID    `json:"-" db:"list_id"`
	CourseId       []uuid.UUID  `json:"-" db:"course_id"`
	Course         []LidCourse  `json:"course" db:"-"`
	StaffID        uuid.UUID    `json:"-" db:"user_id"`
	User           UserDataList `json:"user" db:"-"`
	List           ListData     `json:"list" db:"-"`
	FullName       string       `json:"fullName" db:"full_name"  regex:"login" lenMin:"3" lenMax:"64"`
	PhoneNumber    string       `json:"phoneNumber" db:"phone_number"  regex:"phone" lenMin:"3" lenMax:"16"`
	Location       string       `json:"location" db:"from_location"  regex:"login" lenMin:"0" lenMax:"64"`
	Comment        string       `json:"comment" db:"comment"  regex:"login" lenMin:"0" lenMax:"64"`
	Date           string       `json:"date" db:"created_at"`
	Status         string       `json:"status" db:"status"`
	LastConnection time.Time    `json:"lastConnection" db:"last_connection"`
}
type CreateLid struct {
	ListID         uuid.UUID   `json:"listID" db:"list_id"`
	CourseId       []uuid.UUID `json:"courseId" db:"course_id"`
	StaffId        uuid.UUID   `json:"-" db:"user_id"`
	FullName       string      `json:"fullName" db:"full_name" required:"true" regex:"login" lenMin:"3" lenMax:"64"`
	PhoneNumber    string      `json:"phoneNumber" db:"phone_number" required:"true" regex:"phone" lenMin:"3" lenMax:"16"`
	Location       string      `json:"location" db:"from_location" required:"true"  regex:"login" lenMin:"0" lenMax:"64"`
	Comment        string      `json:"comment" db:"comment"  regex:"login" lenMin:"0" lenMax:"1024"`
	Status         string      `json:"status" db:"status" regex:"login" lenMin:"0" lenMax:"1024"`
	LastConnection time.Time   `json:"lastConnection" db:"last_connection"`
}
type UpdateLid struct {
	ID             uuid.UUID   `json:"-" db:"id"`
	ListID         uuid.UUID   `json:"listID" db:"list_id"`
	CourseId       []uuid.UUID `json:"courseId" db:"course_id"`
	StaffId        uuid.UUID   `json:"-" db:"user_id"`
	FullName       string      `json:"fullName" db:"full_name"  regex:"login" lenMin:"3" lenMax:"64"`
	PhoneNumber    string      `json:"phoneNumber" db:"phone_number" regex:"phone" lenMin:"3" lenMax:"16"`
	Location       string      `json:"location" db:"from_location"  regex:"login" lenMin:"0" lenMax:"64"`
	Comment        string      `json:"comment" db:"comment"  regex:"login" lenMin:"0" lenMax:"1024"`
	Status         string      `json:"status" db:"status" regex:"login" lenMin:"0" lenMax:"1024"`
	LastConnection time.Time   `json:"lastConnection" db:"last_connection"`
}
type MoveLid struct {
	From uuid.UUID `json:"from" db:"from"`
	To   uuid.UUID `json:"to" db:"to"`
}
type MoveList struct {
	From uuid.UUID `json:"from" db:"from"`
	To   uuid.UUID `json:"to" db:"to"`
}
type LidCourse struct {
	CourseID    uuid.UUID `json:"id" db:"-"`
	CourseTitle string    `json:"title" db:"-"`
}
type LidFilter struct {
	SearchKey string
	Course    uuid.UUID
	User      uuid.UUID
	ToDate    time.Time
	EndedDate time.Time
}
