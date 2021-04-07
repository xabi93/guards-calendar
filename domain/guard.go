package domain

import (
	"errors"
	"time"

	"github.com/xabi93/guards-calendar/pkg/id"
)

// different guard slots
const (
	GuardSlotMorning GuardSlot = iota + 1
	GuardSlotAfternoon
	GuardSlotNight
)

// GuardSlot is the guard time range
type GuardSlot int

var guardSlots = map[GuardSlot]string{
	GuardSlotMorning:   "morning",
	GuardSlotAfternoon: "afternoon",
	GuardSlotNight:     "night",
}

func (gs GuardSlot) String() string {
	return guardSlots[gs]
}

// ErrUnknownGuardSlot is returned if the provided guard slot is not defined
var ErrUnknownGuardSlot = errors.New("unknown guard slot")

func (gs GuardSlot) validate() error {
	_, ok := guardSlots[gs]
	if !ok {
		return ErrUnknownGuardSlot
	}
	return nil
}

type GuardDate time.Time

// Before reports whether the guard date day is before the given date
func (gd GuardDate) Before(date time.Time) bool {
	return time.Time(gd).Before(date.Truncate(24 * time.Hour))
}

// String returns the date formated 2006-01-02
func (gd GuardDate) String() string {
	return time.Time(gd).UTC().Format("2006-01-02")
}

// ErrGuardDateBefore is returned when trying to create a guard with already passed date.
var ErrGuardDateBefore = errors.New("guard date cannot be past")

// CreateGuard returns a Guard initialized
func CreateGuard(date GuardDate, slot GuardSlot) (Guard, error) {
	if date.Before(time.Now()) {
		return Guard{}, ErrGuardDateBefore
	}
	if err := slot.validate(); err != nil {
		return Guard{}, err
	}

	return Guard{ID: id.New(), Date: date, Slot: slot}, nil
}

// Guard represents a day time slot for volunteering
type Guard struct {
	ID   id.ID     `json:"id,omitempty"`
	Date GuardDate `json:"date,omitempty"`
	Slot GuardSlot `json:"slot,omitempty"`
}
