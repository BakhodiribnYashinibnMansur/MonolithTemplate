package cronjobs

import (
	"EduCRM/tools/logger"

	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

// var (
// errorNoRowsAffected = errors.New("no rows affected")
// )
type CronJobsDB struct {
	db      *sqlx.DB
	store   *minio.Client
	loggers *logger.Logger
}

func NewCronJobsMethod(db *sqlx.DB, store *minio.Client, loggers *logger.Logger) *CronJobsDB {
	return &CronJobsDB{db: db, store: store, loggers: loggers}
}
