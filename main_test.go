package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenAndServe(t *testing.T) {
  assert := assert.New(t)
  
  t.Run("returns Pepper's score", func  (t *testing.T)  {
    request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
    response := httptest.NewRecorder()

    PlayerServer(response, request)

    got := response.Body.String()
    want := "20"

    assert.Equal(got, want)
  })
}

