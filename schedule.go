package schedule

import (
	"errors"
	"time"
)

// Schedule to store and retrieve appointments
type Schedule struct {
	// active       bool
	rules        []rule
	appointments []schedulable
}

type schedulable interface {
	Datetime() time.Time
}

// ErrInvalidAppointment is a generic error for appoinment related operations
var ErrInvalidAppointment = errors.New("invalid appointment")

// ErrOutOfTimeRange is the error given when try to make an appointmenrt out of the schedule range
// var ErrOutOfDateOrTimeRange = errors.New("The given date or time for the appointment is out of the range defined in the schedule rules")

func (s *Schedule) isApplicableAndValid(appointment schedulable) bool {
	rules := s.getApplicableRulesFor(appointment)
	if len(rules) == 0 {
		return false
	}

	for _, rule := range rules {
		if !rule.isValid(appointment) {
			return false
		}
	}
	return true
}

func (s *Schedule) getApplicableRulesFor(appointment schedulable) []rule {
	rules := []rule{}
	for _, rule := range s.rules {
		if rule.isApplicable(appointment) {
			rules = append(rules, rule)
		}
	}
	return rules
}

// NewAppointment add an appointment to the schedule if possible
func (s *Schedule) NewAppointment(dateTime time.Time) (schedulable, error) {
	newAppointment := appointment{datetime: dateTime}
	if !s.isApplicableAndValid(newAppointment) {
		return nil, ErrInvalidAppointment
	}

	s.appointments = append(s.appointments, newAppointment)
	return newAppointment, nil
}

// NewSpecialAppointment add an appointment to the schedule without caring about the rules
func (s *Schedule) NewSpecialAppointment(dateTime time.Time) (schedulable, error) {
	newAppointment := appointment{datetime: dateTime}

	s.appointments = append(s.appointments, newAppointment)
	return newAppointment, nil
}

// Appointments returns the existing appointments of the schedule
func (s *Schedule) Appointments() []schedulable {
	return s.appointments
}

// func (s *Schedule) FreeAppointments() []schedulable

type appointment struct {
	datetime time.Time
}

func (a appointment) Datetime() time.Time {
	return a.datetime
}

// type Scheduler struct
