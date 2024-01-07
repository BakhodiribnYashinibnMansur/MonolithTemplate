package user

import (
	"EduCRM/model"
	"EduCRM/package/repository/function"
	"EduCRM/tools/logger"
	"database/sql"
	"errors"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type EmployeeReaderDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

func NewEmployeeReaderDB(db *sqlx.DB, loggers *logger.Logger) *EmployeeReaderDB {
	return &EmployeeReaderDB{db: db, loggers: loggers}
}
func (repo *EmployeeReaderDB) GetEmployeeList(role []string, pagination *model.Pagination) (userList []model.Employee, err error) {
	loggers := repo.loggers
	db := repo.db
	err = function.GetListCount(repo.db, loggers, pagination, GetEmployeeCountQuery, []interface{}{pq.Array(role)})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	err = db.Select(&userList, GetEmployeeListByRoleQuery, pq.Array(role), pagination.Limit, pagination.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	return userList, err
}
func (repo *EmployeeReaderDB) GetEmployeeByID(id string, role []string) (user model.Employee, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&user, GetEmployeeByIDQuery, id, pq.Array(role))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		loggers.Error(err)
		return user, err
	}
	return user, err
}
func (repo *EmployeeReaderDB) CheckEmployeeByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var courseID string
	err = db.Get(&courseID, CheckEmployeeByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
		loggers.Error(err)
		return err
	}
	return nil
}
func (repo *EmployeeReaderDB) CheckEmployeeByPhone(phone string) (id string, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&id, CheckEmployeeByPhoneQuery, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, err
		}
		loggers.Error(err)
		return id, err
	}
	return id, nil
}
func (repo *EmployeeReaderDB) GetDropDownEmployee(role []string, filter string, pagination *model.Pagination) (employeeList []model.EmployeeDropDown, err error) {
	loggers := repo.loggers
	db := repo.db
	filterQuery := ``
	number := 1
	var filtersArray []interface{}
	filtersArray = append(filtersArray, pq.Array(role))
	if filter == "" {
		filterQuery = filterQuery + ` `
	}
	if filter != "" {
		number++
		filterQuery = ` AND (full_name iLIKE '%'||$` + strconv.Itoa(number) + `||'%' OR phone_number iLIKE '%'||$` + strconv.Itoa(number) + `||'%')`
		filtersArray = append(filtersArray, filter)
	}
	listCountQuery := GetDropDownEmployeeCountQuery + filterQuery
	query := GetDropDownEmployeeListQuery + filterQuery + ` ORDER BY full_name` + ` LIMIT $` + strconv.Itoa(number+1) + ` OFFSET $` + strconv.Itoa(number+2)
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
func (repo *EmployeeReaderDB) GetEmployeeCount(role []string) (count int64, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&count, GetEmployeeCountByRoleQuery, pq.Array(role))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return count, err
		}
		loggers.Error(err)
		return count, err
	}
	return count, nil
}
