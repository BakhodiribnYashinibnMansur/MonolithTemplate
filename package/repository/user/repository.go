package user

import (
	"EduCRM/model"
	"EduCRM/tools/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	User
	Manager
	Teacher
	Student
	StudentData
	Employee
}
type User struct {
	UserReader
	UserWriter
}
type UserReader interface {
	GetUserList(role string, user *model.Pagination) (userList []model.User, err error)
	GetUserByID(id string) (user model.User, err error)
	SignInUser(user model.SignInUser) (id, role uuid.UUID, err error)
	CheckUserByID(id string) (err error)
	GetUserFullNameById(id string) (string, error)
	GetGroupStudentList(studentIDList []uuid.UUID) (studentList []model.User, err error)
	CheckUserByPhone(phone string) (id string, err error)
	GetUserName(id string) (name string, err error)
	GetUserDataListByIDs(userIDList []uuid.UUID) (userDataList []model.UserDataList, err error)
	GetRoleRelationalCount(id string) (count int64, err error)
	GetDropDownUser(filter string, pagination *model.Pagination) (employeeList []model.UserDropDown, err error)
}
type UserWriter interface {
	CreateUser(user model.CreateUser) (id uuid.UUID, err error)
	UpdateUser(user model.UpdateUser) (err error)
	DeleteUser(id string) (err error)
	UpdateUserPassword(id, password string) error
}
type Manager struct {
	ManagerReader
	ManagerWriter
}
type ManagerReader interface {
	GetManagerList(role string, pagination *model.Pagination) (userList []model.Manager, err error)
	GetManagerByID(id string) (user model.Manager, err error)
	CheckManagerByID(id string) (err error)
	CheckManagerByPhone(phone string) (id string, err error)
}
type ManagerWriter interface {
	CreateManager(user model.CreateManager) (id uuid.UUID, err error)
	UpdateManager(user model.UpdateManager) (err error)
	DeleteManager(id string) (err error)
	UpdateManagerPassword(id, password string) error
}
type Teacher struct {
	TeacherReader
	TeacherWriter
}
type TeacherReader interface {
	GetTeacherList(role string, pagination *model.Pagination) (userList []model.Teacher, err error)
	GetTeacherPageHomeList(role string) (teacherFile []model.TeacherList, err error)
	GetTeacherByID(id string) (user model.Teacher, err error)
	CheckTeacherByID(id string) (err error)
	CheckTeacherByPhone(phone string) (id string, err error)
	GetFilterTeacherList(role string, filter string, pagination *model.Pagination) (teacher []model.TeacherDropDown, err error)
	GetTeacherCount(role string) (count int64, err error)
}
type TeacherWriter interface {
	CreateTeacher(user model.CreateTeacher) (id uuid.UUID, err error)
	UpdateTeacher(user model.UpdateTeacher) (err error)
	DeleteTeacher(id string) (err error)
	UpdateTeacherPassword(id, password string) error
}
type Student struct {
	StudentReader
	StudentWriter
}
type StudentReader interface {
	GetStudentList(role string, pagination *model.Pagination, filter model.StudentFilter) (userList []model.Students, err error)
	GetDropDownStudentList(role string, search string, pagination *model.Pagination) (studentList []model.StudentDropDown, err error)
	GetStudentByID(id string) (user model.Student, err error)
	CheckStudentByID(id string) (err error)
	CheckStudentByPhone(phone string) (id string, err error)
	GetStudentLineChart(role, month string) (count []model.StudentLineChart, err error)
	GetStudentCount(role string) (count int64, err error)
	GetStudentFullDataList() (userList []model.Student, err error)
	GetStudentIDByNameOrPhone(nameOrPhone string) (id []uuid.UUID, err error)
}
type StudentWriter interface {
	CreateStudent(user model.CreateStudent) (id uuid.UUID, err error)
	UpdateStudent(user model.UpdateStudent) (err error)
	DeleteStudent(id string) (err error)
	UpdateStudentPassword(id, password string) error
}
type StudentData struct {
	StudentDataReader
	StudentDataWriter
}
type StudentDataReader interface {
	GetStudentDataByID(id string) (user model.StudentData, err error)
}
type StudentDataWriter interface {
	CreateStudentData(user model.CreateStudent) (id uuid.UUID, err error)
	UpdateStudentData(user model.UpdateStudent) (err error)
	DeleteStudentData(id string) (err error)
}
type Employee struct {
	EmployeeReader
	EmployeeWriter
}
type EmployeeReader interface {
	GetEmployeeList(role []string, pagination *model.Pagination) (userList []model.Employee, err error)
	GetEmployeeByID(id string, role []string) (user model.Employee, err error)
	CheckEmployeeByID(id string) (err error)
	CheckEmployeeByPhone(phone string) (id string, err error)
	GetEmployeeCount(role []string) (count int64, err error)
	GetDropDownEmployee(role []string, filter string, pagination *model.Pagination) (employeeList []model.EmployeeDropDown, err error)
}
type EmployeeWriter interface {
	CreateEmployee(user model.CreateEmployee) (id uuid.UUID, err error)
	UpdateEmployee(user model.UpdateEmployee) (err error)
	DeleteEmployee(id string) (err error)
	UpdateEmployeePassword(id, password string) error
}

func NewUserRepo(db *sqlx.DB, loggers *logger.Logger) *UserRepo {
	return &UserRepo{
		User: User{
			UserReader: NewUserReaderDB(db, loggers),
			UserWriter: NewUserWriterDB(db, loggers),
		},
		Manager: Manager{
			ManagerReader: NewManagerReaderDB(db, loggers),
			ManagerWriter: NewManagerWriterDB(db, loggers),
		},
		Teacher: Teacher{
			TeacherReader: NewTeacherReaderDB(db, loggers),
			TeacherWriter: NewTeacherWriterDB(db, loggers),
		},
		Student: Student{
			StudentReader: NewStudentReaderDB(db, loggers),
			StudentWriter: NewStudentWriterDB(db, loggers),
		},
		Employee: Employee{
			EmployeeReader: NewEmployeeReaderDB(db, loggers),
			EmployeeWriter: NewEmployeeWriterDB(db, loggers),
		},
		StudentData: StudentData{
			StudentDataReader: NewStudentDataReaderDB(db, loggers),
			StudentDataWriter: NewStudentDataWriterDB(db, loggers),
		},
	}
}
