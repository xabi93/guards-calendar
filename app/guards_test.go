package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/xabi93/guards-calendar/app"
	"github.com/xabi93/guards-calendar/domain"
	"github.com/xabi93/guards-calendar/pkg/errors"
)

type GuardsStoreMock struct {
	mock.Mock
}

func (m *GuardsStoreMock) Add(ctx context.Context, guard domain.Guard) error {
	args := m.Called(ctx, guard)
	return args.Error(0)
}

func (m *GuardsStoreMock) Exists(ctx context.Context, day domain.GuardDate, slot domain.GuardSlot) (bool, error) {
	args := m.Called(ctx, day, slot)
	return args.Bool(0), args.Error(1)
}

type GuardsServiceSuite struct {
	suite.Suite

	svc         *app.GuardsService
	guardsStore *GuardsStoreMock

	ctx       context.Context
	guardDate domain.GuardDate
	guardSlot domain.GuardSlot
}

func (s *GuardsServiceSuite) SetupTest() {
	s.guardsStore = &GuardsStoreMock{}
	s.svc = app.NewGuardsService(s.guardsStore)

	s.ctx = context.Background()
	s.guardDate = domain.GuardDate(time.Now())
	s.guardSlot = domain.GuardSlotAfternoon
}

func (s GuardsServiceSuite) TestInvalidParams() {
	_, err := s.svc.Publish(context.Background(), domain.GuardDate(time.Now()), -1)
	s.True(errors.IsWrongInput(err))
}

func (s GuardsServiceSuite) TestAlreadyExists() {
	s.guardsStore.On("Exists", s.ctx, s.guardDate, s.guardSlot).Return(true, nil)
	_, err := s.svc.Publish(s.ctx, s.guardDate, s.guardSlot)
	s.True(errors.IsConflict(err))
	s.ErrorIs(err, &app.GuardAlreadyExistsError{Date: s.guardDate, Slot: s.guardSlot})
}

func TestGuardsService(t *testing.T) {
	suite.Run(t, new(GuardsServiceSuite))
}
