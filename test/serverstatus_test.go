package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerStatusGetStatus(t *testing.T) {
	t.Run("Should return 200 when everything is Ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := newContext(req, rec)

		c.setRouteTester("/server/status", nil, nil)

		if assert.NoError(t, serverstatus.GetStatus(c.c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
