package domain_test

import (
	"fmt"
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

func TestGuardDate(t *testing.T) {
	year := 2021
	month := time.December
	day := 17

	require.Equal(t, fmt.Sprintf("%d-%d-%d", year, month, day), domain.GuardDate(time.Date(year, month, day, 5, 22, 3, 1, time.UTC)).String())
}

func TestCreateGuard(t *testing.T) {
	now := time.Now()
	t.Run("invalid", func(t *testing.T) {
		for _, tc := range []struct {
			date        domain.GuardDate
			slot        domain.GuardSlot
			expectedErr error
		}{
			{
				date:        domain.GuardDate(now.AddDate(0, 0, -1)),
				slot:        domain.GuardSlotMorning,
				expectedErr: domain.ErrGuardDateBefore,
			},
			{
				date:        domain.GuardDate(now),
				expectedErr: domain.ErrUnknownGuardSlot,
			},
		} {
			_, err := domain.CreateGuard(tc.date, tc.slot)
			require.Equal(t, tc.expectedErr, err)
		}
	})

	t.Run("valid", func(t *testing.T) {
		for _, tc := range []struct {
			date domain.GuardDate
			slot domain.GuardSlot
		}{
			{
				date: domain.GuardDate(now),
				slot: domain.GuardSlotMorning,
			},
			{
				date: domain.GuardDate(now),
				slot: domain.GuardSlotAfternoon,
			},
			{
				date: domain.GuardDate(now),
				slot: domain.GuardSlotNight,
			},
		} {
			guard, err := domain.CreateGuard(tc.date, tc.slot)
			require.NoError(t, err)

			require.NoError(t, guard.ID.Validate())
			require.Equal(t, tc.date, guard.Date)
			require.Equal(t, tc.slot, guard.Slot)
		}
	})
}
