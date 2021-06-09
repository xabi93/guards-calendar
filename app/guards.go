package app

import (
	"context"
	"fmt"

	"github.com/xabi93/guards-calendar/domain"
	"github.com/xabi93/guards-calendar/pkg/errors"
)

// GuardAlreadyExistsError is returned when trying to create a guard already exists for a date and slot
type GuardAlreadyExistsError struct {
	Date domain.GuardDate
	Slot domain.GuardSlot
}

func (err GuardAlreadyExistsError) Error() string {
	return fmt.Sprintf("guard for day %s and slot %s already exists", err.Date, err.Slot)
}

// NewGuardsService returns an initialized GuardService
func NewGuardsService(guards GuardsStore) *GuardsService {
	return &GuardsService{guards}
}

// GuardsService provides functions for guards
type GuardsService struct {
	guards GuardsStore
}

// Publish creates a new guard given a date and a guard slot if does not already exists
func (g GuardsService) Publish(ctx context.Context, date domain.GuardDate, slot domain.GuardSlot) (domain.Guard, error) {
	guard, err := domain.CreateGuard(date, slot)
	if err != nil {
		return domain.Guard{}, errors.WrapWrongInput(err, "publishing guard")
	}

	exists, err := g.guards.Exists(ctx, guard.Date, guard.Slot)
	if err != nil {
		return domain.Guard{}, err
	}
	if exists {
		return domain.Guard{}, errors.WrapConflict(
			GuardAlreadyExistsError{Date: date, Slot: slot},
			"publishing guard",
		)
	}

	if err := g.guards.Add(ctx, guard); err != nil {
		return domain.Guard{}, err
	}

	return guard, nil
}
