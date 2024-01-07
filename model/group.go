package model

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	Title           string     `json:"title" db:"title"  required:"true"`
	Description     string     `json:"description" db:"description"`
	CourseID        uuid.UUID  `json:"-" db:"course_id"`
	Course          CourseData `json:"course" db:"-"`
	TeacherID       uuid.UUID  `json:"-" db:"teacher_id"`
	Teacher         UserData   `json:"teacher" db:"-"`
	EduDays         string     `json:"eduDays" db:"edu_days" required:"true"`
	RoomID          uuid.UUID  `json:"-" db:"room_id"`
	Room            RoomData   `json:"room" db:"-"`
	Price           int        `json:"price" db:"price"`
	PhotoName       string     `json:"-" db:"photo"`
	Photo           File       `json:"photo" db:"-"`
	Status          bool       `json:"status" db:"status"`
	LessonStartTime string     `json:"lessonStartTime" db:"lesson_start_time"`
	LessonEndTime   string     `json:"lessonEndTime" db:"lesson_end_time"`
	StartDate       string     `json:"startDate" db:"start_date"`
	EndDate         string     `json:"endDate" db:"end_date"`
	Comment         string     `json:"comment" db:"comment"`
	StudentCount    int64      `json:"studentCount" db:"student_count"`
}
type CreateGroup struct {
	Title           string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	Description     string    `json:"description" db:"description"  lenMin:"0" lenMax:"1024" regex:"login"`
	CourseID        uuid.UUID `json:"courseID" db:"course_id" required:"true"`
	TeacherID       uuid.UUID `json:"teacherID" db:"teacher_id" required:"true"`
	EduDays         string    `json:"eduDays" db:"edu_days" required:"true" lenMin:"0" lenMax:"32" regex:"login"`
	RoomID          uuid.UUID `json:"roomID" db:"room_id"`
	Photo           string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	Price           int       `json:"price" db:"price"  required:"true" amountMin:"100000" amountMax:"100000000" regex:"login"`
	Duration        string    `json:"-" db:"-"  `
	Status          bool      `json:"status" db:"status" `
	LessonStartTime string    `json:"lessonStartTime" db:"lesson_start_time"  lenMin:"0" lenMax:"16" regex:"login"`
	LessonEndTime   string    `json:"lessonEndTime" db:"lesson_end_time"`
	StartDate       time.Time `json:"startDate" db:"start_date" required:"true"`
	EndDate         time.Time `json:"endDate" db:"end_date" required:"true"`
	Comment         string    `json:"comment" db:"comment" lenMin:"0" lenMax:"2048" regex:"login"`
}
type UpdateGroup struct {
	ID              uuid.UUID `json:"-" db:"id"`
	Title           string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64"`
	Description     string    `json:"description" db:"description"  lenMin:"0" lenMax:"1024"`
	CourseID        uuid.UUID `json:"courseID" db:"course_id" required:"true"`
	TeacherID       uuid.UUID `json:"teacherID" db:"teacher_id" required:"true"`
	EduDays         string    `json:"eduDays" db:"edu_days" required:"true" lenMin:"0" lenMax:"32"`
	RoomID          uuid.UUID `json:"roomID" db:"room_id"`
	Price           int       `json:"price" db:"price"  required:"true" amountMin:"100000" amountMax:"100000000"`
	Photo           string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	Status          bool      `json:"status" db:"status" `
	LessonStartTime string    `json:"lessonStartTime" db:"lesson_start_time"  lenMin:"0" lenMax:"16"`
	LessonEndTime   string    `json:"lessonEndTime" db:"lesson_end_time" lenMin:"0" lenMax:"16"`
	StartDate       time.Time `json:"startDate" db:"start_date" required:"true"`
	EndDate         time.Time `json:"endDate" db:"end_date" required:"true"`
	Comment         string    `json:"comment" db:"comment" lenMin:"0" lenMax:"2048"`
}
type GroupDropDown struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	PhotoName string    `json:"-" db:"photo"`
	Photo     File      `json:"photo" db:"-"`
}
type SearchGroupDropDown struct {
	SearchTitleGroup string
}
type GroupFilter struct {
	Status    string    `json:"status" db:"status"`
	TeacherId uuid.UUID `json:"teacher_id" db:"teacher_id"`
	CourseId  uuid.UUID `json:"course_id" db:"course_id"`
}
