package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/xabi93/guards-calendar/domain"
)

func TestGuardSlot(t *testing.T) {
	for slot, str := range map[domain.GuardSlot]string{
		domain.GuardSlotMorning:   "morning",
		domain.GuardSlotAfternoon: "afternoon",
		domain.GuardSlotNight:     "night",
		0:                         "",
	} {
		require.Equal(t, str, slot.String())
	}
}

func TestGuard(t *testing.T) {
	t.Run("Create guard", func(t *testing.T) {
		date := time.Now()
		slot := domain.GuardSlotMorning

		guard := domain.CreateGuard(date, slot)

		require.NoError(t, guard.ID.Validate())
		require.Equal(t, date, guard.Date)
		require.Equal(t, slot, guard.Slot)
	})
}
