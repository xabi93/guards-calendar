package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xabi93/guards-calendar/api"
)

func TestHealthHandler(t *testing.T) {
	h := api.HealthHandler()

	rw := httptest.NewRecorder()
	h(rw, nil)

	require.Equal(t, http.StatusOK, rw.Result().StatusCode)
}
