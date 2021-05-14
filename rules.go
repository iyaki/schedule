package schedule

import "time"

type rule interface {
	isApplicable(schedulable) bool
	isValid(schedulable) bool
}

type baseRule struct {
	weekDay        time.Weekday
	startDate      time.Time
	endDate        time.Time
	dailyStartTime time.Time
	dailyEndTime   time.Time
}

func (r baseRule) isApplicable(appointment schedulable) bool {
	return r.weekDay == appointment.Datetime().Weekday() &&
		isDateInRange(appointment.Datetime(), r.startDate, r.endDate) &&
		isTimeInRange(appointment.Datetime(), r.dailyStartTime, r.dailyEndTime)
}

type dateTimeRule struct {
	baseRule
}

func (r dateTimeRule) isValid(appointment schedulable) bool {
	return true
}

var _ rule = dateTimeRule{}
