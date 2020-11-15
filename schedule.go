package schedule

import (
	"time"
)

// Schedule to store and retrieve appointments
type Schedule struct {
	appointments []appointment
}

type appointment struct {
	Datetime time.Time
}

// NewAppointment add an appointment to the schedule if possible
func (s *Schedule) NewAppointment(dateTime time.Time) appointment {
	app := appointment{Datetime: dateTime}
	s.appointments = append(s.appointments, app)
	return app
}

// Appointments returns the existing appointments of the schedule
func (s *Schedule) Appointments() []appointment {
	return s.appointments
}
