package errors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xabi93/guards-calendar/pkg/errors"
)

func TestErrorTypes(t *testing.T) {
	err := fmt.Errorf("some error")
	for _, tc := range []struct {
		err    error
		isFunc func(err error) bool
	}{
		{
			err:    errors.WrapWrongInput(err, "wrapping"),
			isFunc: errors.IsWrongInput,
		},
		{
			err:    errors.WrapNotFound(err, "wrapping"),
			isFunc: errors.IsNotFound,
		},
		{
			err:    errors.WrapConflict(err, "wrapping"),
			isFunc: errors.IsConflict,
		},
		{
			err:    errors.NewConflict("some error"),
			isFunc: errors.IsConflict,
		},
	} {
		require.True(t, tc.isFunc(tc.err))
	}
}
