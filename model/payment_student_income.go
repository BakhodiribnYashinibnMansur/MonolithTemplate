package model
import (
	"time"
	"github.com/google/uuid"
)
type PaymentIncome struct {
	Id            uuid.UUID        `json:"id" db:"id"`
	GroupID       uuid.UUID        `json:"-" db:"group_id"`
	StudentID     uuid.UUID        `json:"-" db:"student_id"`
	Student       UserData         `json:"student" db:"-"`
	Amount        int64            `json:"amount" db:"amount"`
	PaymentTypeID uuid.UUID        `json:"-" db:"payment_type_id"`
	PaymentType   PaymentTypeData  `json:"paymentType" db:"-"`
	TeacherID     uuid.UUID        `json:"-" db:"teacher_id"`
	Teacher       UserData         `json:"teacher" db:"-"`
	Group         StudentGroupList `json:"group" db:"-"`
	StaffID       uuid.UUID        `json:"-" db:"staff_id"`
	Staff         UserData         `json:"staff" db:"-"`
	Comment       string           `json:"comment" db:"comment"`
	CreatedAt     string           `json:"createdAt" db:"created_at"`
	UpdatedAt     string           `json:"updatedAt" db:"updated_at"`
}
type CreatePaymentIncome struct {
	StudentID     uuid.UUID `json:"studentID" db:"student_id"`
	Amount        int64     `json:"amount" db:"amount"  required:"true" regex:"login" amountMin:"1000" amountMax:"100000000000"`
	GroupID       uuid.UUID `json:"groupID" db:"group_id"`
	PaymentTypeID uuid.UUID `json:"paymentTypeId" db:"payment_type_id"`
	Comment       string    `json:"comment" db:"comment" lenMin:"0" lenMax:"1024" regex:"login" `
	StaffID       uuid.UUID `json:"-" db:"staff_id"`
}
type UpdatePaymentIncome struct {
	Id            uuid.UUID `json:"-" db:"id"`
	StudentID     uuid.UUID `json:"studentID" db:"student_id"`
	Amount        int64     `json:"amount" db:"amount"  required:"true" regex:"login" amountMin:"1000" amountMax:"100000000000"`
	GroupID       uuid.UUID `json:"groupID" db:"group_id"`
	PaymentTypeID uuid.UUID `json:"paymentTypeID" db:"payment_type_id"`
	Comment       string    `json:"comment" db:"comment" lenMin:"0" lenMax:"1024" regex:"login" `
	StaffID       uuid.UUID `json:"-" db:"staff_id"`
}
type StudentIncomeFilter struct {
	TypeID              uuid.UUID
	StudentNameOrPhone  string
	StudentIDs          []uuid.UUID
	TeacherID           uuid.UUID
	StaffID             uuid.UUID
	TeacherGroupsIDList []uuid.UUID
	GroupID             uuid.UUID
	FromDate            time.Time
	ToDate              time.Time
}
type PaymentIncomeStatistic struct {
	LineChart            []PaymentIncomeLineChart `json:"lineChart" db:"-"`
	TotalAmountLineChart int                      `json:"totalAmountLineChart" db:"-"`
	PieChart             []PaymentIncomePieChart  `json:"pieChart" db:"-"`
	TotalAmountPieChart  int                      `json:"totalAmountPieChart" db:"-"`
}
type PaymentIncomeLineChart struct {
	Target string `json:"target" db:"target"`
	Month  string `json:"month" db:"month"`
	Year   string `json:"year" db:"year"`
	Amount int    `json:"amount" db:"amount"`
}
type PaymentIncomePieChart struct {
	Title      string  `json:"title" db:"full_name"`
	Percentage float64 `json:"percentage" db:"Percentage"`
	Amount     int64   `json:"amount" db:"amount"`
}
