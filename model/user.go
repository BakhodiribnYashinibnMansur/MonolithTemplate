package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID               uuid.UUID          `json:"id"  db:"id"`
	FullName         string             `json:"fullName" db:"full_name"`
	BirthdayDate     string             `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber      string             `json:"phoneNumber" db:"phone_number"`
	RoleID           uuid.UUID          `json:"roleID" db:"role_id"`
	Role             string             `json:"role" db:"-"`
	Photo            string             `json:"photo" db:"photo"`
	PhotoLink        string             `json:"photoLink" db:"-"`
	AddedDate        string             `json:"addedDate" db:"added_date"`
	Total            int64              `json:"-" db:"total"`
	GroupList        []Group            `json:"groupList" db:"-"`
	CourseList       []Course           `json:"courseList" db:"-"`
	StudentGroup     []Group            `json:"studentGroup" db:"-"`
	StudentGroupList []StudentGroupList `json:"studentGroupList" db:"-"`
}
type CreateUser struct {
	ID           uuid.UUID `json:"-"  db:"id"`
	FullName     string    `json:"fullName" db:"full_name" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string    `json:"phoneNumber" db:"phone_number" required:"true" lenMin:"0" lenMax:"16" regex:"phone"`
	RoleID       string    `json:"roleID" db:"role_id" lenMin:"0" lenMax:"64" regex:"login"`
	RoleTitle    string    `json:"roleTitle" db:"-" lenMin:"0" lenMax:"64" regex:"login"`
	Password     string    `json:"password" db:"password" required:"true" lenMin:"0" lenMax:"64"`
	Photo        string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	AddedDate    time.Time `json:"addedDate" db:"added_date" required:"true"`
}
type UpdateUser struct {
	ID           uuid.UUID `json:"-"  db:"id"`
	FullName     string    `json:"fullName" db:"full_name" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string    `json:"phoneNumber" db:"phone_number" required:"true" lenMin:"0" lenMax:"16" regex:"phone"`
	RoleID       string    `json:"roleID" db:"role_id" lenMin:"0" lenMax:"64" regex:"login"`
	RoleTitle    string    `json:"roleTitle" db:"-" lenMin:"0" lenMax:"64" regex:"login"`
	Photo        string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	AddedDate    time.Time `json:"addedDate" db:"added_date" required:"true"`
}
type SignInUser struct {
	PhoneNumber string `json:"phoneNumber" db:"phone_number" default:"+998901234567" required:"true"  lenMin:"0" lenMax:"16" regex:"phone"`
	Password    string `json:"password" db:"password" default:"EduCRM$007Boss" `
}
type SignInUserResponse struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Role uuid.UUID `json:"role" db:"role_id"`
}
type UserPassword struct {
	Password string `json:"password"`
}
type StudentGroupList struct {
	ID         uuid.UUID `json:"id" db:"group_id"`
	GroupTitle string    `json:"title" db:"title"`
}
type CreateSuperAdmin struct {
	Token       string `json:"token" db:"token" required:"true" lenMin:"0" lenMax:"64"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number" default:"+998901234567" required:"true"  lenMin:"0" lenMax:"16" regex:"phone"`
	Password    string `json:"password" db:"password" default:"EduCRM$007Boss" `
}
