package cronjobs

import (
	"EduCRM/tools/logger"

	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

type CronJobs struct {
	CronJobsMethod
}
type CronJobsMethod interface {
}

func NewCronJobs(db *sqlx.DB, minio *minio.Client, loggers *logger.Logger) *CronJobs {
	cronjobs := &CronJobs{
		CronJobsMethod: NewCronJobsMethod(db, minio, loggers),
	}
	// RunCronJobs("* * * * *", loggers, func() error {})
	return cronjobs
}
