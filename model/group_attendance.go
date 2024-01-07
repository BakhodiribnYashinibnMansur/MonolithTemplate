package model

import (
	"github.com/google/uuid"
	"time"
)

type ModifyGroupSchedule struct {
	ID                uuid.UUID `json:"-" db:"id"`
	GroupID           uuid.UUID `json:"-" db:"group_id"`
	Date              time.Time `json:"date" db:"date"`
	ThemeTitle        string    `json:"themeTitle" db:"theme" lenMin:"5" lenMax:"1024"`
	LessonOrderNumber int       `json:"-" db:"lesson_order_number" amountMin:"0" amountMax:"100"`
}
type GroupSchedule struct {
	ID       string     `json:"id" db:"id"`
	Month    []int64    `json:"month" db:"-"`
	Schedule []Schedule `json:"schedule" db:"-"`
}
type GroupStudentAttendance struct {
	GroupID           uuid.UUID           `json:"groupId" db:"group_id"`
	StudentAttendance []StudentAttendance `json:"studentAttendance" db:"-"`
}
type StudentAttendance struct {
	StudentID uuid.UUID `json:"studentId" db:"student_id"`
	Absent    bool      `json:"absent" db:"absent"`
	Homework  bool      `json:"homework" db:"homework"`
}
type Schedule struct {
	ID                uuid.UUID `json:"id" db:"id"`
	ThemeTitle        string    `json:"themeTitle" db:"theme"`
	LessonOrderNumber int       `json:"lessonOrderNumber" db:"lesson_order_number"`
	Date              string    `json:"date" db:"date"`
}
type ModifyGroupAttendance struct {
	GroupID           uuid.UUID           `json:"-" db:"-"`
	GroupScheduleID   uuid.UUID           `json:"groupScheduleID" db:"group_id"`
	StudentAttendance []StudentAttendance `json:"studentAttendance" db:"-"`
}
type StudentAttendanceDB struct {
	LessonOrderNumber int64     `json:"-" db:"lesson_order_number"`
	Date              string    `db:"date"`
	GroupScheduleID   uuid.UUID `json:"groupScheduleID" db:"group_schedule_id"`
	StudentID         uuid.UUID `json:"studentId" db:"student_id"`
	Attendance        bool      `json:"absent" db:"absent"`
	Homework          bool      `json:"homework" db:"homework"`
}
type GroupAttendanceUI struct {
	StudentID         uuid.UUID             `json:"studentId" db:"student_id"`
	StudentFullName   string                `json:"studentFullName" db:"-"`
	StudentAttendance []StudentAttendanceUI `json:"studentAttendance" db:"-"`
}
type StudentAttendanceUI struct {
	LessonOrderNumber int64  `json:"lessonOrderNumber" db:"lesson_order_number"`
	Date              string `json:"date" db:"date"`
	Attendance        bool   `json:"attendance" db:"attendance"`
	Homework          bool   `json:"homework" db:"homework"`
}
