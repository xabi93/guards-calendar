package domain

import (
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

func (gs GuardSlot) String() string {
	switch gs {
	case GuardSlotMorning:
		return "morning"
	case GuardSlotAfternoon:
		return "afternoon"
	case GuardSlotNight:
		return "night"
	}
	return ""
}

// CreateGuard returns a Guard initialized
func CreateGuard(date time.Time, slot GuardSlot) Guard {
	return Guard{ID: id.New(), Date: date, Slot: slot}
}

// Guard represents a day time slot for volunteering
type Guard struct {
	ID   id.ID     `json:"id,omitempty"`
	Date time.Time `json:"date,omitempty"`
	Slot GuardSlot `json:"slot,omitempty"`
}
