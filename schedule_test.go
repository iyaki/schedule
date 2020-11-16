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

func TestNewAppointment(t *testing.T) {
	t.Run("Add new valid appointment to a schedule", func(t *testing.T) {
		startTime := time.Date(0, 0, 0, 8, 0, 0, 0, time.UTC)
		endTime := time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC)
		appointmentDateTime := time.Date(0, 0, 0, 9, 30, 0, 0, time.UTC)
		schedule := Schedule{
			StartTime: startTime,
			EndTime:   endTime,
		}

		newAppointment, err := schedule.NewAppointment(appointmentDateTime)

		appointments := schedule.Appointments()
		appointmentsAmmount := len(appointments)
		gettedAppointment := appointments[0]

		assertNoError(t, err)
		assertEqual(t, appointmentDateTime, gettedAppointment.Datetime())
		assertEqual(t, newAppointment, gettedAppointment)
		assertEqual(t, appointmentsAmmount, 1)
	})

	t.Run("Add new appointment to a schedule out of hour range", func(t *testing.T) {
		startTime := time.Date(0, 0, 0, 8, 0, 0, 0, time.UTC)
		endTime := time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC)
		appointmentDateTime := time.Date(0, 0, 0, 6, 30, 0, 0, time.UTC)
		schedule := Schedule{
			StartTime: startTime,
			EndTime:   endTime,
		}

		newAppointment, err := schedule.NewAppointment(appointmentDateTime)

		appointmentsAmmount := len(schedule.Appointments())

		assertError(t, err, ErrOutOfTimeRange)
		assertEqual(t, newAppointment, nil)
		assertEqual(t, appointmentsAmmount, 0)
	})
}

func assertEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	if expected != got {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("getted an error but didn't want one")
	}
}
