package model

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID               uuid.UUID       `json:"id"  db:"id"`
	FullName         string          `json:"fullName" db:"full_name"`
	BirthdayDate     string          `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber      string          `json:"phoneNumber" db:"phone_number"`
	RoleID           uuid.UUID       `json:"-" db:"role_id"`
	Role             RoleData        `json:"role" db:"-"`
	PhotoName        string          `json:"-" db:"photo"`
	Photo            File            `json:"photo" db:"-"`
	Amount           int64           `json:"amount" db:"amount"`
	AddedDate        string          `json:"addedDate" db:"added_date"`
	Comment          string          `json:"commit" db:"commit"`
	CellularPhone    string          `json:"cellularPhone" db:"cellular_phone"`
	StudentGivenID   string          `json:"studentGivenID" db:"student_given_id"`
	ParentsPhone     string          `json:"parentsPhone" db:"parents_phone"`
	Email            string          `json:"email" db:"email"`
	TelegramNickName string          `json:"telegramNickName" db:"telegram_nick_name"`
	Location         string          `json:"location" db:"location"`
	Passport         string          `json:"passport" db:"password"`
	Discount         int64           `json:"discount" db:"discount"`
	Group            []Group         `json:"studentGroup" db:"-"`
	Payment          []PaymentIncome `json:"payment" db:"-"`
	Tag              StudentTag      `json:"tag" db:"-"`
	Status           StudentStatus   `json:"status" db:"-"`
	TagId            uuid.UUID       `json:"-" db:"student_tag"`
	StatusId         uuid.UUID       `json:"-" db:"student_status"`
}
type Students struct {
	ID           uuid.UUID          `json:"id"  db:"id"`
	FullName     string             `json:"fullName" db:"full_name"`
	BirthdayDate string             `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string             `json:"phoneNumber" db:"phone_number"`
	Amount       int64              `json:"amount" db:"amount"`
	RoleID       uuid.UUID          `json:"-" db:"role_id"`
	Role         RoleData           `json:"role" db:"-"`
	PhotoName    string             `json:"-" db:"photo"`
	Photo        File               `json:"photo" db:"-"`
	AddedDate    string             `json:"addedDate" db:"added_date"`
	Student      []StudentGroupList `json:"studentGroup" db:"-"`
	Tag          StudentTag         `json:"tag" db:"-" `
	Status       StudentStatus      `json:"status" db:"-"`
	TagId        uuid.UUID          `json:"-" db:"student_tag" `
	StatusId     uuid.UUID          `json:"-" db:"student_status"`
	// GroupCount   int64              `json:"groupCount" db:"group_count"`
}
type CreateStudent struct {
	ID               uuid.UUID   `json:"-"  db:"id"`
	FullName         string      `json:"fullName" db:"full_name" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate     time.Time   `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber      string      `json:"phoneNumber" db:"phone_number" required:"true" lenMin:"0" lenMax:"16" regex:"phone"`
	Password         string      `json:"password" db:"password"  lenMin:"8" lenMax:"64"`
	RoleID           string      `json:"-" db:"role_id"`
	RoleTitle        string      `json:"-" db:"-"`
	Photo            string      `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	AddedDate        time.Time   `json:"addedDate" db:"added_date" required:"true"`
	GroupID          []uuid.UUID `json:"groupID" db:"group_id"`
	Commit           string      `json:"commit" db:"commit" lenMin:"0" lenMax:"1024" regex:"login" `
	CellularPhone    string      `json:"cellularPhone" db:"cellular_phone" lenMin:"0" lenMax:"16" regex:"phone"`
	StudentGivenID   string      `json:"studentGivenID" db:"student_given_id" lenMin:"0" lenMax:"64" regex:"login"`
	ParentsPhone     string      `json:"parentsPhone" db:"parents_phone" lenMin:"0" lenMax:"16" regex:"phone"`
	Email            string      `json:"email" db:"email" lenMin:"0" lenMax:"32" regex:"email"`
	TelegramNickName string      `json:"telegramNickName" db:"telegram_nick_name" lenMin:"0" lenMax:"32" regex:"login"`
	Location         string      `json:"location" db:"location" lenMin:"0" lenMax:"512" regex:"login"`
	Passport         string      `json:"passport" db:"password" lenMin:"0" lenMax:"64" regex:"login"`
	TagID            uuid.UUID   `json:"tagID" db:"tag"`
	Discount         int64       `json:"discount" db:"discount" amountMin:"0" amountMax:"100"`
	StatusID         uuid.UUID   `json:"statusID" db:"status"`
}
type UpdateStudent struct {
	ID               uuid.UUID   `json:"-"  db:"id"`
	FullName         string      `json:"fullName" db:"full_name" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate     time.Time   `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber      string      `json:"phoneNumber" db:"phone_number" required:"true" lenMin:"0" lenMax:"16" regex:"phone"`
	RoleID           string      `json:"-" db:"role_id"`
	RoleTitle        string      `json:"-" db:"-"`
	Photo            string      `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	AddedDate        time.Time   `json:"addedDate" db:"added_date" required:"true"`
	GroupID          []uuid.UUID `json:"groupID" db:"group_id"`
	Commit           string      `json:"commit" db:"commit" lenMin:"0" lenMax:"1024" regex:"login" `
	CellularPhone    string      `json:"cellularPhone" db:"cellular_phone" lenMin:"0" lenMax:"16" regex:"login"`
	StudentGivenID   string      `json:"studentGivenID" db:"student_given_id" lenMin:"0" lenMax:"64" regex:"login"`
	ParentsPhone     string      `json:"parentsPhone" db:"parents_phone" lenMin:"0" lenMax:"16" regex:"regex"`
	Email            string      `json:"email" db:"email" lenMin:"0" lenMax:"32" regex:"email"`
	TelegramNickName string      `json:"telegramNickName" db:"telegram_nick_name" lenMin:"0" lenMax:"32" regex:"login"`
	Location         string      `json:"location" db:"location" lenMin:"0" lenMax:"512" regex:"login"`
	Passport         string      `json:"passport" db:"password" lenMin:"0" lenMax:"64" regex:"login"`
	TagID            uuid.UUID   `json:"tagID" db:"tag" lenMin:"0"`
	Discount         int64       `json:"discount" db:"discount" amountMin:"0" amountMax:"100"`
	StatusID         uuid.UUID   `json:"statusID" db:"status"`
}
type StudentData struct {
	StudentID        uuid.UUID `json:"-"  db:"id"`
	Commit           string    `json:"commit" db:"commit" lenMin:"0" lenMax:"1024" regex:"login" `
	CellularPhone    string    `json:"cellularPhone" db:"cellular_phone" lenMin:"0" lenMax:"16" regex:"login"`
	StudentGivenID   string    `json:"studentID" db:"student_given_id" lenMin:"0" lenMax:"64" regex:"login"`
	ParentsPhone     string    `json:"parentsPhone" db:"parents_phone" lenMin:"0" lenMax:"16" regex:"login"`
	Email            string    `json:"email" db:"email" lenMin:"0" lenMax:"32" regex:"email"`
	TelegramNickName string    `json:"telegramNickName" db:"telegram_nick_name" lenMin:"0" lenMax:"32" regex:"login"`
	Location         string    `json:"location" db:"location" lenMin:"0" lenMax:"512" regex:"login"`
	Passport         string    `json:"passport" db:"passport" lenMin:"0" lenMax:"64" regex:"login"`
	Tag              uuid.UUID `json:"tag" db:"student_tag"`
	Discount         int64     `json:"discount" db:"discount" amountMin:"0" amountMax:"100"`
	Status           uuid.UUID `json:"status" db:"student_status"`
}
type StudentDropDown struct {
	ID          uuid.UUID `json:"id"  db:"id"`
	FullName    string    `json:"fullName" db:"full_name"`
	PhoneNumber string    `json:"-" db:"phone_number"`
}
type StudentFilter struct {
	SearchKey   string      `json:"search_key" db:"-"`
	Course      uuid.UUID   `json:"course" db:"-"`
	CourseUsers []uuid.UUID `json:"courseUsers" db:"-"`
	Group       uuid.UUID   `json:"group" db:"-"`
	GroupUsers  []uuid.UUID `json:"groupUsers" db:"-"`
	Status      string      `json:"status" db:"-"`
	Tag         string      `json:"tag" db:"-"`
}
