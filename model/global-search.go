package model

type GlobalOutPutData struct {
	User   []UserDropDown   `json:"user" db:"-"`
	Group  []GroupDropDown  `json:"group" db:"-"`
	Course []CourseDropDown `json:"course" db:"-"`
}
