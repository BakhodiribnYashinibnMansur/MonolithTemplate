package model

import "github.com/google/uuid"

type Statistic struct {
	CourseCount        int64                   `json:"courseCount" db:"id"`
	GroupCount         int64                   `json:"groupCount" db:"id"`
	StudentCount       int64                   `json:"studentCount" db:"id"`
	TeacherCount       int64                   `json:"teacherCount" db:"id"`
	EmployeeCount      int64                   `json:"employeeCount" db:"id"`
	StudentLineChart   []StudentLineChart      `json:"studentLineChart" db:"-"`
	TotalAmountStudent int                     `json:"totalAmount" db:"-"`
	Teacher            []TeacherList           `json:"teacherList" db:"-"`
	IncomePieChart     []PaymentIncomePieChart `json:"incomePieChart" db:"-"`
	ExpensePieChart    []ExpensePieChart       `json:"expensePieChart" db:"-"`
	TotalAmountIncome  int                     `json:"totalAmountIncome" db:"-"`
	TotalAmountExpense int                     `json:"totalAmountExpense" db:"-"`
}
type ExpensePieChart struct {
	Title      string `json:"title" db:"title"`
	Percentage int    `json:"percentage" db:"Percentage"`
	Amount     int    `json:"amount" db:"amount"`
}
type TeacherList struct {
	ID        uuid.UUID `json:"teacherID" db:"id"`
	FullName  string    `json:"fullName" db:"full_name"`
	PhotoName string    `json:"-" db:"photo"`
	Photo     File      `json:"photo" db:"-"`
}
type StudentLineChart struct {
	Month        string `json:"month" db:"month"`
	Year         string `json:"year" db:"year"`
	StudentCount int    `json:"studentCount" db:"student_count"`
}
