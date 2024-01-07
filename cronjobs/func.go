package cronjobs

import (
	"EduCRM/tools/logger"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	errInvalidCronExpression = errors.New("invalid cron value number is ")
	errInvalidCronFormat     = errors.New("invalid cron format")
)

type cronFunc = func() error
type CronJobsExpression struct {
	MinuteNumerator       int64 `min:"0" max:"59"`
	MinuteDenominator     int64 `min:"0" max:"59"`
	HourNumerator         int64 `min:"0" max:"23"`
	HourDenominator       int64 `min:"0" max:"23"`
	DayOfMonthNumerator   int64 `min:"1" max:"31"`
	DayOfMonthDenominator int64 `min:"1" max:"31"`
	MonthNumerator        int64 `min:"1" max:"12"`
	MonthDenominator      int64 `min:"1" max:"12"`
	DayOfWeekNumerator    int64 `min:"0" max:"6"`
	DayOfWeekDenominator  int64 `min:"0" max:"6"`
}

var CronJobsExpressionList = []int64{}

func RunCronJobs(schedule string, loggers *logger.Logger, cronFunc cronFunc) {
	var cron CronJobsExpression
	cron, err := getCronJobsExpression(schedule, loggers)
	if err != nil {
		loggers.Error(err)
		return
	}
	var wg sync.WaitGroup
	triggerMinute := time.NewTicker(time.Minute)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			currentDate := <-triggerMinute.C
			isCron := map[bool]bool{}
			for i := 0; i < len(CronJobsExpressionList); i = i + 2 {
				ans := cron.getCronDate(CronJobsExpressionList[i], CronJobsExpressionList[i+1], getDatePart(currentDate, int64(i)))
				isCron[ans] = ans
			}
			if ok := isCron[true]; ok && len(isCron) == 1 {
				err := cronFunc()
				if err != nil {
					loggers.Error(err)
					return
				}
			}
		}
	}()
	wg.Wait()
}
func getDatePart(time time.Time, target int64) int64 {
	if target == 0 {
		return int64(time.Minute())
	} else if target == 2 {
		return int64(time.Hour())
	} else if target == 4 {
		return int64(time.Day())
	} else if target == 6 {
		return int64(time.Month())
	} else if target == 8 {
		return int64(time.Weekday())
	}
	return 0
}
func (cron CronJobsExpression) getCronDate(numerator, denominator, target int64) bool {
	isNumerator := false
	isDenominator := false
	if numerator == 0 && denominator == 0 {
		return true
	}
	if numerator == 0 && target%denominator == 0 {
		isNumerator = true
		isDenominator = true
	}
	if denominator > 0 && numerator > 0 {
		if target >= numerator {
			isNumerator = true
		}
		if (target-numerator)%denominator == 0 {
			isDenominator = true
		}
	}
	return isNumerator && isDenominator
}
func getCronJobsExpression(expression string, loggers *logger.Logger) (cron CronJobsExpression, err error) {
	parts := strings.Fields(expression)
	if len(parts) != 5 {
		return CronJobsExpression{}, errInvalidCronFormat
	}
	cron.MinuteNumerator, cron.MinuteDenominator, err = getCronNumber(parts[0], 1)
	if err != nil {
		loggers.Error(err)
		return CronJobsExpression{}, err
	}
	cron.HourNumerator, cron.HourDenominator, err = getCronNumber(parts[1], 2)
	if err != nil {
		loggers.Error(err)
		return CronJobsExpression{}, err
	}
	cron.DayOfMonthNumerator, cron.DayOfMonthDenominator, err = getCronNumber(parts[2], 3)
	if err != nil {
		loggers.Error(err)
		return CronJobsExpression{}, err
	}
	cron.MonthNumerator, cron.MonthDenominator, err = getCronNumber(parts[3], 4)
	if err != nil {
		loggers.Error(err)
		return CronJobsExpression{}, err
	}
	cron.DayOfWeekNumerator, cron.DayOfWeekDenominator, err = getCronNumber(parts[4], 5)
	if err != nil {
		loggers.Error(err)
		return CronJobsExpression{}, err
	}
	err = validateCron(cron)
	if err != nil {
		loggers.Error(err)
		return CronJobsExpression{}, err
	}
	CronJobsExpressionList = append(CronJobsExpressionList, cron.MinuteNumerator, cron.MinuteDenominator, cron.HourNumerator, cron.HourDenominator, cron.DayOfMonthNumerator, cron.DayOfMonthDenominator, cron.MonthNumerator, cron.MonthDenominator, cron.DayOfWeekNumerator, cron.DayOfWeekDenominator)
	return cron, nil
}
func getCronNumber(number string, order int64) (numerator int64, denominator int64, err error) {
	if number == "*" {
		return 0, 0, nil
	}
	if strings.Contains(number, "/") {
		numberParts := strings.Split(number, "/")
		if len(numberParts) != 2 {
			return 0, 0, errors.New(errInvalidCronExpression.Error() + strconv.Itoa(int(order)))
		}
		if numberParts[0] == "*" {
			numerator = 0
		} else {
			numerator, err = strconv.ParseInt(numberParts[0], 10, 64)
			if err != nil {
				return 0, 0, err
			}
		}
		denominator, err := strconv.ParseInt(numberParts[1], 10, 64)
		if err != nil {
			return 0, 0, err
		}
		return numerator, denominator, nil
	}
	numerator, err = strconv.ParseInt(number, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return numerator, 1, nil
}
func validateCron(model CronJobsExpression) error {
	minTag := "min"
	maxTag := "max"
	structType := reflect.TypeOf(model)
	structValue := reflect.ValueOf(model)
	for i := 0; i < structType.NumField(); i++ {
		fieldName := structType.Field(i).Name
		fieldValue := reflect.Indirect(structValue).FieldByName(fieldName)
		amountMinTagValue, amountMinTagExist := structType.Field(i).Tag.
			Lookup(minTag)
		amountMaxTagValue, amountMaxTagExist := structType.Field(i).Tag.
			Lookup(maxTag)
		if amountMinTagExist && amountMaxTagExist {
			amountMinTagIntValue, err := strconv.Atoi(amountMinTagValue)
			if err != nil {
				return errors.New(err.Error())
			}
			amountMaxTagIntValue, err := strconv.Atoi(amountMaxTagValue)
			if err != nil {
				return errors.New(err.Error())
			}
			if amountMinTagIntValue-1 >= int(fieldValue.Int()) || int(fieldValue.Int()) >= amountMaxTagIntValue+1 {
				return errors.New(fieldName + " must be " + strconv.Itoa(amountMinTagIntValue) + " and " + strconv.Itoa(amountMaxTagIntValue) + " length")
			}
			return nil
		}
	}
	return nil
}
