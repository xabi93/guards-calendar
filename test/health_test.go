package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xabi93/guards-calendar/cmd/guards-calendar/server"
)

func TestHealth(t *testing.T) {
	s := httptest.NewServer(server.New())

	resp, err := s.Client().Get(s.URL + "/health")

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
