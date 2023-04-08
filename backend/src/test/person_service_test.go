package test

import (
	"context"
	"github.com/stretchr/testify/require"
	"net.http"
	"testing"

	"bityagi/service"
	"github.com/stretchr/testify/assert"
)

func TestShowPerson(t *testing.T) {
	s := service.NewMyPersonApiService()
	ctx := context.Background()
	resp, err := s.ShowPerson(ctx)
	require.NoError(t, err, "error should be nil")
	require.NoNil(t, resp, "response should not be nil")
	
	assert.Equal(t, http.StatusOK, resp.Code, "status code should be 200")
}