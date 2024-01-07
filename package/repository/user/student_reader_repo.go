package user

import (
	"EduCRM/model"
	"EduCRM/package/repository/function"
	"EduCRM/tools/logger"
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type StudentReaderDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

func NewStudentReaderDB(db *sqlx.DB, loggers *logger.Logger) *StudentReaderDB {
	return &StudentReaderDB{db: db, loggers: loggers}
}
func (repo *StudentReaderDB) GetStudentList(role string, pagination *model.Pagination, filter model.StudentFilter) (userList []model.Students, err error) {
	loggers := repo.loggers
	db := repo.db
	filterQuery := ``
	number := 1
	var filtersArray []interface{}
	filtersArray = append(filtersArray, role)
	if filter.SearchKey != "" {
		number++
		filterQuery = filterQuery + ` AND (TRIM(BOTH ' ' FROM crm_user.full_name) iLIKE '%'||$` + strconv.Itoa(number) + `||'%' OR crm_user.phone_number iLIKE '%'||$` + strconv.Itoa(number) + `||'%')`
		filtersArray = append(filtersArray, strings.ReplaceAll(filter.SearchKey, " ", ""))
	}
	if filter.Tag != "" {
		number++
		filterQuery = filterQuery + ` AND student_data.student_tag =$` + strconv.Itoa(number)
		filtersArray = append(filtersArray, filter.Tag)
	}
	if filter.Status != "" {
		number++
		filterQuery = filterQuery + ` AND student_data.student_status =$` + strconv.Itoa(number)
		filtersArray = append(filtersArray, filter.Status)
	}
	if filter.Group != uuid.Nil {
		number++
		filterQuery = filterQuery + ` AND crm_user.id =Any($` + strconv.Itoa(number) + `)`
		filtersArray = append(filtersArray, pq.Array(filter.GroupUsers))
	}
	if filter.Course != uuid.Nil {
		number++
		filterQuery = filterQuery + ` AND crm_user.id = ANY($` + strconv.Itoa(number) + `)`
		filtersArray = append(filtersArray, pq.Array(filter.CourseUsers))
	}
	countFilterQuery := GetStudentListCountQuery + filterQuery
	query := GetStudentListQuery + filterQuery + `ORDER BY crm_user.created_at DESC ` + `  LIMIT $` + strconv.Itoa(number+1) + ` OFFSET $` + strconv.Itoa(number+2)
	err = function.GetListCount(repo.db, loggers, pagination, countFilterQuery, filtersArray)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	filtersArray = append(filtersArray, pagination.Limit)
	filtersArray = append(filtersArray, pagination.Offset)
	err = db.Select(&userList, query, filtersArray...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	return userList, err
}
func (repo *StudentReaderDB) GetStudentByID(id string) (user model.Student, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&user, GetStudentByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		loggers.Error(err)
		return user, err
	}
	return user, err
}
func (repo *StudentReaderDB) CheckStudentByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var courseID string
	err = db.Get(&courseID, CheckStudentByIDQuery, id)
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
func (repo *StudentReaderDB) GetGroupStudentList(studentIDList []uuid.UUID) (studentList []model.Students, err error) {
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
func (repo *StudentReaderDB) CheckStudentByPhone(phone string) (id string, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&id, CheckStudentByPhoneQuery, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, err
		}
		loggers.Error(err)
		return id, err
	}
	return id, nil
}
func (repo *StudentReaderDB) GetDropDownStudentList(role string, search string, pagination *model.Pagination) (student []model.StudentDropDown, err error) {
	loggers := repo.loggers
	db := repo.db
	filterQuery := ``
	number := 1
	var filterArray []interface{}
	filterArray = append(filterArray, role)
	if search != "" {
		number++
		filterQuery = ` AND (TRIM(BOTH ' ' FROM full_name) iLIKE '%'||$` + strconv.Itoa(number) + `||'%' OR phone_number iLIKE '%'||$` + strconv.Itoa(number) + `||'%')`
		filterArray = append(filterArray, strings.ReplaceAll(search, " ", ""))
	}
	listCountQuery := GetSearchStudentCountQuery + filterQuery
	query := GetSearchStudentListQuery + filterQuery + ` ORDER BY full_name` + ` LIMIT $` + strconv.Itoa(number+1) + ` OFFSET $` + strconv.Itoa(number+2)
	err = function.GetListCount(repo.db, loggers, pagination, listCountQuery, filterArray)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return student, err
		}
		loggers.Error(err)
		return student, err
	}
	filterArray = append(filterArray, pagination.Limit)
	filterArray = append(filterArray, pagination.Offset)
	err = db.Select(&student, query, filterArray...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return student, err
		}
		loggers.Error(err)
		return student, err
	}
	return student, nil
}
func (repo *StudentReaderDB) GetStudentCount(role string) (count int64, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&count, GetStudentCountQuery, role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return count, err
		}
		loggers.Error(err)
		return count, err
	}
	return count, nil
}
func (repo *StudentReaderDB) GetStudentLineChart(role, month string) (count []model.StudentLineChart, err error) {
	loggers := repo.loggers
	db := repo.db
	NumberMonth, err := strconv.Atoi(month)
	if err != nil {
		return count, err
	}
	NumberMonth = NumberMonth - 1
	err = db.Select(&count, GetStudentCountLineChartQuery, role, strconv.Itoa(NumberMonth)+" month")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return count, err
		}
		loggers.Error(err)
		return count, err
	}
	return count, nil
}
func (repo *StudentReaderDB) GetStudentFullDataList() (userList []model.Student, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&userList, GetStudentFullDataListQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	return userList, nil
}
func (repo *StudentReaderDB) GetStudentIDByNameOrPhone(nameOrPhone string) (ids []uuid.UUID, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Select(&ids, GetStudentIDByNameOrPhoneQuery, nameOrPhone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ids, err
		}
		loggers.Error(err)
		return ids, err
	}
	repo.loggers.Info(ids)
	return ids, err
}
