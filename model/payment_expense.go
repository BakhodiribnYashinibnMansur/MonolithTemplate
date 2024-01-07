package model

import (
	"time"

	"github.com/google/uuid"
)

type CreatePaymentExpense struct {
	Title            string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"256" regex:"login"`
	Description      string    `json:"description" db:"description" lenMin:"0" lenMax:"1024" regex:"login"`
	Amount           int       `json:"amount" db:"amount" required:"true" regex:"login" amountMin:"1000" amountMax:"100000000000"`
	PaymentTypeId    uuid.UUID `json:"paymentTypeId" db:"payment_type_id"`
	PaymentCatalogId uuid.UUID `json:"paymentCatalogId" db:"payment_catalog_id"`
	StaffId          uuid.UUID `json:"userId" db:"staff_id"`
}
type UpdatePaymentExpense struct {
	Id               uuid.UUID `json:"id" db:"id"`
	Title            string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"256" regex:"login"`
	Description      string    `json:"description" db:"description" lenMin:"0" lenMax:"1024" regex:"login"`
	Amount           int       `json:"amount" db:"amount" required:"true" regex:"login" amountMin:"1000" amountMax:"100000000000"`
	PaymentTypeId    uuid.UUID `json:"paymentTypeId" db:"payment_type_id" `
	PaymentCatalogId uuid.UUID `json:"paymentCatalogId" db:"payment_catalog_id" `
	StaffId          uuid.UUID `json:"userId" db:"staff_id"`
}
type PaymentExpense struct {
	Id                  uuid.UUID          `json:"id" db:"id"`
	Title               string             `json:"title" db:"title"`
	PaymentTypeTitle    string             `json:"-" db:"payment_type_title"`
	PaymentCatalogTitle string             `json:"-" db:"payment_catalog_title"`
	Description         string             `json:"description" db:"description"`
	Amount              int                `json:"amount" db:"amount"`
	PaymentTypeId       uuid.UUID          `json:"-" db:"payment_type_id"`
	PaymentType         PaymentTypeData    `json:"paymentType" db:"-"`
	PaymentCatalogId    uuid.UUID          `json:"-" db:"payment_catalog_id"`
	PaymentCatalog      PaymentCatalogData `json:"paymentCatalog" db:"-"`
	StaffId             uuid.UUID          `json:"-" db:"staff_id"`
	Staff               UserData           `json:"user" db:"-"`
	CreatedAt           string             `json:"createdAt" db:"created_at"`
	UpdatedAt           string             `json:"updatedAt" db:"updated_at"`
}
type ExpenseFilter struct {
	FromDate       time.Time
	ToDate         time.Time
	PaymentType    uuid.UUID
	PaymentCatalog uuid.UUID
	StaffID        uuid.UUID
	SearchKey      string
}
type PaymentExpensesStatistic struct {
	LineChart            []PaymentExpensesLineChart `json:"lineChart" db:"-"`
	TotalAmountLineChart int                        `json:"totalAmountLineChart" db:"-"`
	PieChart             []PaymentExpensePieChart   `json:"pieChart" db:"-"`
	TotalAmountPieChart  int                        `json:"totalAmountPieChart" db:"-"`
}
type PaymentExpensesLineChart struct {
	Month  string `json:"month" db:"month"`
	Target string `json:"target" db:"target"`
	Year   string `json:"year" db:"year"`
	Amount int    `json:"amount" db:"amount"`
}
type PaymentExpensePieChart struct {
	Title      string  `json:"title" db:"title"`
	Percentage float64 `json:"percentage" db:"Percentage"`
	Amount     int     `json:"amount" db:"amount"`
}
