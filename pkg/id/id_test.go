package id_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/xabi93/guards-calendar/pkg/id"
)

func TestID(t *testing.T) {
	t.Run("valid new id", func(t *testing.T) {
		id := id.New()
		_, err := uuid.Parse(string(id))
		require.NoError(t, err)
	})

	t.Run("validate invalid id", func(t *testing.T) {
		require.Error(t, id.ID("asdasda").Validate())
	})

	t.Run("validate valid id", func(t *testing.T) {
		require.NoError(t, id.New().Validate())
	})
}
