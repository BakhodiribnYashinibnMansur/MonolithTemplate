package user

import (
	"EduCRM/model"
	"EduCRM/package/repository/function"
	"EduCRM/tools/logger"
	"database/sql"
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var (
	errorSignInUser = errors.New("  Username or Password is Incorrect")
)

type UserReaderDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

func NewUserReaderDB(db *sqlx.DB, loggers *logger.Logger) *UserReaderDB {
	return &UserReaderDB{db: db, loggers: loggers}
}
func (repo *UserReaderDB) GetUserList(role string, pagination *model.Pagination) (userList []model.User, err error) {
	loggers := repo.loggers
	db := repo.db
	// 	err = db.Select(&userList, GetUserListAllQuery, pagination.Limit, pagination.Offset)
	// 	if err != nil {
	// 		if errors.Is(err, sql.ErrNoRows) {
	// 			return userList, err
	// 		}
	// 		loggers.Error(err)
	// 		return userList, err
	// 	}
	// 	loggers.Error(err)
	// 	return userList, err
	// }
	err = function.GetListCount(repo.db, loggers, pagination, GetUserCountByRoleIDQuery, []interface{}{role})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	err = db.Select(&userList, GetUserListByRoleQuery, role, pagination.Limit, pagination.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	return userList, err
}
func (repo *UserReaderDB) GetUserFullNameById(id string) (roll string, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&roll, GetStudentFullNameByIdQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return roll, err
		}
		loggers.Error(err)
		return roll, err
	}
	return roll, err
}
func (repo *UserReaderDB) GetUserByID(id string) (user model.User, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&user, GetUserByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		loggers.Error(err)
		return user, err
	}
	return user, err
}
func (repo *UserReaderDB) SignInUser(user model.SignInUser) (id, role uuid.UUID, err error) {
	var result model.SignInUserResponse
	loggers := repo.loggers
	err = repo.db.Get(&result, SignInUserQuery, user.PhoneNumber, user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, role, err
		}
		loggers.Error(err)
		return id, role, errorSignInUser
	}
	return result.ID, result.Role, nil
}
func (repo *UserReaderDB) CheckUserByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var courseID string
	err = db.Get(&courseID, CheckUserByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	if courseID != id {
		return errors.New("invalid course id")
	}
	return nil
}
func (repo *UserReaderDB) GetGroupStudentList(studentIDList []uuid.UUID) (studentList []model.User, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&studentList, GetGroupStudentListQuery, studentIDList)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return studentList, err
		}
		loggers.Error(err)
		return studentList, err
	}
	return studentList, nil
}
func (repo *UserReaderDB) CheckUserByPhone(phone string) (id string, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&id, CheckUserByPhoneQuery, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, err
		}
		loggers.Error(err)
		return id, err
	}
	return id, nil
}
func (repo *UserReaderDB) GetUserName(id string) (name string, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&name, GetStudentFullNameByIdQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return name, err
		}
		loggers.Error(err)
		return name, err
	}
	return name, nil
}
func (repo *UserReaderDB) GetUserDataListByIDs(userIDList []uuid.UUID) (userDataList []model.UserDataList, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&userDataList, GetUserDataListByUserIDsList, pq.Array(userIDList))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userDataList, err
		}
		loggers.Error(err)
		return userDataList, err
	}
	return userDataList, nil
}
func (repo *UserReaderDB) GetRoleRelationalCount(id string) (count int64, err error) {
	err = repo.db.Get(GetRoleUserRelationalCountQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return count, err
		}
		repo.loggers.Error(err)
		return count, err
	}
	return count, err
}
func (repo *UserReaderDB) GetDropDownUser(filter string, pagination *model.Pagination) (employeeList []model.UserDropDown, err error) {
	loggers := repo.loggers
	db := repo.db
	filterQuery := ``
	number := 0
	var filtersArray []interface{}
	if filter != "" {
		number++
		filterQuery = filterQuery + ` AND (full_name iLIKE '%'||$` + strconv.Itoa(number) + `||'%' OR phone_number iLIKE '%'||$` + strconv.Itoa(number) + `||'%')`
		filtersArray = append(filtersArray, filter)
	}
	listCountQuery := GetUserCountQuery + filterQuery
	query := GetUserDropDownListQuery + filterQuery + ` ORDER BY full_name` + ` LIMIT $` + strconv.Itoa(number+1) + ` OFFSET $` + strconv.Itoa(number+2)
	err = function.GetListCount(repo.db, loggers, pagination, listCountQuery, filtersArray)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return employeeList, err
		}
		loggers.Error(err)
		return employeeList, err
	}
	filtersArray = append(filtersArray, pagination.Limit)
	filtersArray = append(filtersArray, pagination.Offset)
	err = db.Select(&employeeList, query, filtersArray...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return employeeList, err
		}
		loggers.Error(err)
		return employeeList, err
	}
	return employeeList, nil
}
