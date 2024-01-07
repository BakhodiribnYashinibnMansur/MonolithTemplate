package model

import (
	"github.com/google/uuid"
	"time"
)

type Debtor struct {
	StudentID          uuid.UUID `json:"-" db:"student_id"`
	StudentFullName    string    `json:"-" db:"student_full_name"`
	Student            UserData  `json:"student"`
	StudentPhoneNumber string    `json:"-" db:"student_phone_number"`
	GroupID            uuid.UUID `json:"-" db:"group_id"`
	GroupTitle         string    `json:"-" db:"group_title"`
	Group              GroupData `json:"group"`
	TeacherID          uuid.UUID `json:"-" db:"teacher_id"`
	Teacher            UserData  `json:"teacher"`
	Amount             int64     `json:"amount" db:"amount"`
	Total              int64     `json:"-" db:"total_amount"`
}
type DebtorFilter struct {
	StudentKey string
	StudentIDs []uuid.UUID `json:"-" db:"student_id"`
	GroupID    uuid.UUID   `json:"-" db:"group_id"`
	TeacherID  uuid.UUID   `json:"-" db:"teacher_id"`
	FromDate   time.Time   `json:"-" db:"from_date"`
	ToDate     time.Time   `json:"-" db:"to_date"`
	FromPrice  int64       `json:"-" db:"from_price"`
	ToPrice    int64       `json:"-" db:"to_price"`
}
