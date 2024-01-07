package model

import "github.com/google/uuid"

type TitleIDList struct {
	ID    string `db:"id"`
	Title string `db:"title"`
}
type File struct {
	FileName string `json:"fileName" db:"-"`
	FileLink string `json:"fileLink" db:"-"`
}
type RoleData struct {
	RoleID    uuid.UUID `json:"roleID" db:"-"`
	RoleTitle string    `json:"roleTitle" db:"-"`
}
type UserData struct {
	UserID          uuid.UUID `json:"userID" db:"-"`
	UserFullName    string    `json:"userFullName" db:"-"`
	UserPhoneNumber string    `json:"userPhoneNumber,omitempty" db:"-"`
}
type CourseGroup struct {
	GroupCount int64   `json:"groupCount" db:"-"`
	GroupList  []Group `json:"groupList" db:"-"`
}
type GroupData struct {
	GroupID    uuid.UUID `json:"groupID" db:"-"`
	GroupTitle string    `json:"groupTitle" db:"-"`
}
type CourseData struct {
	CourseID    uuid.UUID `json:"courseID" db:"-"`
	CourseTitle string    `json:"courseTitle" db:"-"`
}
type RoomData struct {
	RoomID    uuid.UUID `json:"roomID" db:"-"`
	RoomTitle string    `json:"roomTitle" db:"-"`
}
type BoardData struct {
	BoardID    uuid.UUID `json:"boardID" db:"-"`
	BoardTitle string    `json:"-" db:"-"`
}
type ListData struct {
	ListID    uuid.UUID `json:"listID" db:"-"`
	ListTitle string    `json:"ListTitle" db:"-"`
}
type LidData struct {
	LidID       uuid.UUID `json:"lidID" db:"-"`
	LidFullName string    `json:"-" db:"-"`
}
type PaymentTypeData struct {
	PaymentTypeID    uuid.UUID `json:"paymentTypeID" db:"-"`
	PaymentTypeTitle string    `json:"paymentTypeTitle" db:"-"`
}
type PaymentCatalogData struct {
	PaymentCatalogID    uuid.UUID `json:"paymentCatalogID" db:"-"`
	PaymentCatalogTitle string    `json:"paymentCatalogTitle" db:"-"`
}

//	type StudentGroupIDData struct {
//		StudentGroupListFullData []StudentGroupList `json:"studentGroupList" db:"-"`
//	}
//
//	type StudentGroupListData struct {
//		StudentGroupList []Group `json:"studentGroupListFullData" db:"-"`
//	}
type TeacherGroupData struct {
	GroupList  []Group `json:"groupList" db:"-"`
	GroupCount int64   `json:"groupCount" db:"-"`
}
type UserDataList struct {
	UserID       uuid.UUID `json:"userID" db:"id"`
	UserFullName string    `json:"userFullName" db:"full_name"`
}
