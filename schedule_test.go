package schedule

import (
	"testing"
	"time"
)

func TestAppointments(t *testing.T) {
	testAppointment := appointment{}
	schedule := Schedule{appointments: []schedulable{testAppointment}}

	got := schedule.Appointments()[0]

	assertEqual(t, testAppointment, got)
}

func newTestSchedule(baseDateTime time.Time) Schedule {
	schedule := Schedule{}

	schedule.rules = append(schedule.rules, newTestRule(baseDateTime))
	
	return schedule
}

func TestNewAppointment(t *testing.T) {
	baseDateTime := time.Date(2021, time.May, 13, 14, 5, 0, 0, time.UTC)

	t.Run("Add new valid appointment to a schedule", func(t *testing.T) {
		schedule := newTestSchedule(baseDateTime)

		newAppointment, err := schedule.NewAppointment(baseDateTime)

		appointments := schedule.Appointments()
		appointmentsAmmount := len(appointments)
		gettedAppointment := appointments[0]

		assertNoError(t, err)
		assertEqual(t, baseDateTime, gettedAppointment.Datetime())
		assertEqual(t, newAppointment, gettedAppointment)
		assertEqual(t, 1, appointmentsAmmount)
	})

	t.Run("Add new appointment to a schedule out of hour range", func(t *testing.T) {
		schedule := newTestSchedule(baseDateTime)

		newAppointment, err := schedule.NewAppointment(baseDateTime.Add(10 * time.Hour))

		assertError(t, err, ErrInvalidAppointment)
		assertEqual(t, nil, newAppointment)
		assertEqual(t, 0, len(schedule.Appointments()))
	})
}
