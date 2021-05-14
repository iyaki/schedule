package schedule

import "time"

func isDateInRange(date time.Time, startDate time.Time, endDate time.Time) bool {
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, endDate.Location())

	return isDateTimeInRange(date, startDate, endDate)
}

func isTimeInRange(t time.Time, startTime time.Time, endTime time.Time) bool {
	startTime = time.Date(t.Year(), t.Month(), t.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), startTime.Nanosecond(), startTime.Location())
	endTime = time.Date(t.Year(), t.Month(), t.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), endTime.Nanosecond(), startTime.Location())

	return isDateTimeInRange(t, startTime, endTime)
}

func isDateTimeInRange(dateTime time.Time, startDateTime time.Time, endDateTime time.Time) bool {
	return (startDateTime.Before(dateTime) ||
		startDateTime.Equal(dateTime)) &&
		(endDateTime.After(dateTime) ||
			endDateTime.Equal(dateTime))
}
