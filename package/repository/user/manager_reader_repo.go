package user

import (
	"EduCRM/model"
	"EduCRM/package/repository/function"
	"EduCRM/tools/logger"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type ManagerReaderDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

func NewManagerReaderDB(db *sqlx.DB, loggers *logger.Logger) *ManagerReaderDB {
	return &ManagerReaderDB{db: db, loggers: loggers}
}
func (repo *ManagerReaderDB) GetManagerList(role string, pagination *model.Pagination) (userList []model.Manager, err error) {
	loggers := repo.loggers
	db := repo.db
	err = function.GetListCount(repo.db, loggers, pagination, GetManagerCountQuery, []interface{}{role})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	err = db.Select(&userList, GetManagerListByRoleQuery, role, pagination.Limit, pagination.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userList, err
		}
		loggers.Error(err)
		return userList, err
	}
	return userList, err
}
func (repo *ManagerReaderDB) GetManagerByID(id string) (user model.Manager, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&user, GetManagerByIDQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		loggers.Error(err)
		return user, err
	}
	return user, err
}
func (repo *ManagerReaderDB) CheckManagerByID(id string) (err error) {
	loggers := repo.loggers
	db := repo.db
	var courseID string
	err = db.Get(&courseID, CheckManagerByIDQuery, id)
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
func (repo *ManagerReaderDB) CheckManagerByPhone(phone string) (id string, err error) {
	loggers := repo.loggers
	db := repo.db
	err = db.Get(&id, CheckManagerByPhoneQuery, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, err
		}
		loggers.Error(err)
		return id, err
	}
	return id, nil
}
