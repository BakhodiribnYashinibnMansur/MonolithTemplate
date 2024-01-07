package user

import (
	"EduCRM/model"
	"EduCRM/package/repository/function"
	"EduCRM/tools/logger"
	"database/sql"
	"errors"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type TeacherReaderDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

func NewTeacherReaderDB(db *sqlx.DB, loggers *logger.Logger) *TeacherReaderDB {
	return &TeacherReaderDB{db: db, loggers: loggers}
}
func (repo *TeacherReaderDB) GetTeacherList(role string, pagination *model.Pagination) (userList []model.Teacher, err error) {
	loggers := repo.loggers
	db := repo.db
	err = function.GetListCount(repo.db, loggers, pagination, GetTeacherCountQuery, []interface{}{role})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	err = db.Select(&userList, GetTeacherListByRoleQuery, role, pagination.Limit, pagination.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	return userList, err
}
func (repo *TeacherReaderDB) GetTeacherByID(id string) (user model.Teacher, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&user, GetTeacherByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		loggers.Error(err)
		return user, err
	}
	return user, err
}
func (repo *TeacherReaderDB) CheckTeacherByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var courseID string
	err = db.Get(&courseID, CheckTeacherByIDQuery, id)
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
func (repo *TeacherReaderDB) CheckTeacherByPhone(phone string) (id string, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&id, CheckTeacherByPhoneQuery, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, err
		}
		loggers.Error(err)
		return id, err
	}
	return id, nil
}
func (repo *TeacherReaderDB) GetFilterTeacherList(role string, filter string, pagination *model.Pagination) (teacher []model.TeacherDropDown, err error) {
	loggers := repo.loggers
	db := repo.db
	filterQuery := ``
	number := 1
	var filtersArray []interface{}
	filtersArray = append(filtersArray, role)
	if filter == "" {
		filterQuery = filterQuery + ` `
	} else if filter != "" {
		number++
		filterQuery = ` AND (full_name iLIKE '%'||$` + strconv.Itoa(number) + `||'%' OR phone_number iLIKE '%'||$` + strconv.Itoa(number) + `||'%')`
		filtersArray = append(filtersArray, filter)
	}
	listCountQuery := GetFilterTeacherCountQuery + filterQuery
	query := GetFilterTeacherListQuery + filterQuery + ` ORDER BY full_name` + ` LIMIT $` + strconv.Itoa(number+1) + ` OFFSET $` + strconv.Itoa(number+2)
	err = function.GetListCount(repo.db, loggers, pagination, listCountQuery, filtersArray)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return teacher, err
		}
		loggers.Error(err)
		return teacher, err
	}
	filtersArray = append(filtersArray, pagination.Limit)
	filtersArray = append(filtersArray, pagination.Offset)
	err = db.Select(&teacher, query, filtersArray...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return teacher, err
		}
		loggers.Error(err)
		return teacher, err
	}
	return teacher, nil
}
func (repo *TeacherReaderDB) GetTeacherCount(role string) (count int64, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&count, GetTeacherCountQuery, role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return count, err
		}
		loggers.Error(err)
		return count, err
	}
	return count, nil
}
func (repo *TeacherReaderDB) GetTeacherPageHomeList(role string) (teacherFile []model.TeacherList, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&teacherFile, GetTeacherStatisticListQuery, role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return teacherFile, err
		}
		loggers.Error(err)
		return teacherFile, err
	}
	return teacherFile, err
}
