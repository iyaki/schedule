package schedule

import (
	"errors"
	"time"
)

// Schedule to store and retrieve appointments
type Schedule struct {
	StartTime    time.Time
	EndTime      time.Time
	appointments []schedulable
}

type schedulable interface {
	Datetime() time.Time
}

// ErrOutOfTimeRange is the error given when try to make an appointmenrt out of the schedule range
var ErrOutOfTimeRange = errors.New("The given time for the appointment is out of the range defined in the schedule")

// NewAppointment add an appointment to the schedule if possible
func (s *Schedule) NewAppointment(dateTime time.Time) (schedulable, error) {
	if dateTime.Format("15:04:05") < s.StartTime.Format("15:04:05") ||
		dateTime.Format("15:04:05") > s.EndTime.Format("15:04:05") {
		return nil, ErrOutOfTimeRange
	}

	newAppointment := appointment{datetime: dateTime}
	s.appointments = append(s.appointments, newAppointment)
	return newAppointment, nil
}

// Appointments returns the existing appointments of the schedule
func (s *Schedule) Appointments() []schedulable {
	return s.appointments
}

type appointment struct {
	datetime time.Time
}

func (a appointment) Datetime() time.Time {
	return a.datetime
}
