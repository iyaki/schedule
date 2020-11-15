package schedule

import (
	"testing"
	"time"
)

func TestAppointments(t *testing.T) {
	testAppointment := appointment{}
	schedule := Schedule{appointments: []appointment{testAppointment}}

	got := schedule.Appointments()[0]

	assertEqual(t, testAppointment, got)
}

func TestNewAppointment(t *testing.T) {
	schedule := Schedule{}

	newAppointment := schedule.NewAppointment(time.Now())
	gettedAppointment := schedule.Appointments()[0]

	assertEqual(t, newAppointment, gettedAppointment)
}

func assertEqual(t *testing.T, expected, got interface{}) {
	t.Helper()
	if expected != got {
		t.Errorf("expected %v but got %v", expected, got)
	}
}
