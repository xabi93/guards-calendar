package app

import (
	"context"

	"github.com/xabi93/guards-calendar/domain"
)

// GuardsStore is responsible of guards persistance
type GuardsStore interface {
	// Add persists the given guard
	Add(ctx context.Context, guard domain.Guard) error
	// Exists checks if a guard already exists by date and slot
	Exists(ctx context.Context, day domain.GuardDate, slot domain.GuardSlot) (bool, error)
}
