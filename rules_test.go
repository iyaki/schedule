package schedule

import (
	"testing"
	"time"
)

func newTestRule(baseDateTime time.Time) rule {
	weekDay := baseDateTime.Weekday()
	startDate := baseDateTime.AddDate(0, -14, 0)
	endDate := baseDateTime.AddDate(0, 14, 0)
	dailyStartTime := baseDateTime.Add(-time.Hour)
	dailyEndTime := baseDateTime.Add(time.Hour)

	return dateTimeRule{
		basicRule: basicRule{
			weekDay:        weekDay,
			startDate:      startDate,
			endDate:        endDate,
			dailyStartTime: dailyStartTime,
			dailyEndTime:   dailyEndTime,
		},
	}
}

func TestDateTimeRule(t *testing.T) {
	baseDateTime := time.Date(2021, time.May, 13, 14, 5, 0, 0, time.UTC)

	rule := newTestRule(baseDateTime)

	t.Run("Valid appointment same day", func(t *testing.T) {
		appointment := appointment{datetime: baseDateTime}

		assertEqual(t, true, rule.isApplicable(appointment))
		assertEqual(t, true, rule.isValid(appointment))
	})

	t.Run("Valid appointment last week", func(t *testing.T) {
		appointment := appointment{datetime: baseDateTime.AddDate(0, 0, -7)}

		assertEqual(t, true, rule.isApplicable(appointment))
		assertEqual(t, true, rule.isValid(appointment))
	})

	t.Run("Valid appointment next week", func(t *testing.T) {
		appointment := appointment{datetime: baseDateTime.AddDate(0, 0, 7)}

		assertEqual(t, true, rule.isApplicable(appointment))
		assertEqual(t, true, rule.isValid(appointment))
	})

	t.Run("Invalid appointment because different weekDay", func(t *testing.T) {
		appointment := appointment{datetime: baseDateTime.AddDate(0, 1, 0)}

		assertEqual(t, false, rule.isApplicable(appointment))
		assertEqual(t, true, rule.isValid(appointment))
	})

	t.Run("Invalid appointment because sooner than start date", func(t *testing.T) {
		appointment := appointment{datetime: baseDateTime.AddDate(0, -21, 0)}

		assertEqual(t, false, rule.isApplicable(appointment))
		assertEqual(t, true, rule.isValid(appointment))
	})

	t.Run("Invalid appointment because later than end date", func(t *testing.T) {
		appointment := appointment{datetime: baseDateTime.AddDate(0, 21, 0)}

		assertEqual(t, false, rule.isApplicable(appointment))
		assertEqual(t, true, rule.isValid(appointment))
	})

	t.Run("Invalid appointment because sooner than daily start time", func(t *testing.T) {
		appointment := appointment{datetime: baseDateTime.Add(-2 * time.Hour)}

		assertEqual(t, false, rule.isApplicable(appointment))
		assertEqual(t, true, rule.isValid(appointment))
	})

	t.Run("Invalid appointment because later than daily end time", func(t *testing.T) {
		appointment := appointment{datetime: baseDateTime.Add(2 * time.Hour)}

		assertEqual(t, false, rule.isApplicable(appointment))
		assertEqual(t, true, rule.isValid(appointment))
	})
}
