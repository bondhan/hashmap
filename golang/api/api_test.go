package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInit(t *testing.T) {
	req, err := http.NewRequest("GET", "/init", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Init)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "init", rr.Body.String())
}

func TestPut(t *testing.T) {
	req, err := http.NewRequest("GET", "/put", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Put)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "put", rr.Body.String())
}

func TestGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/get", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Get)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "get", rr.Body.String())
}

func TestRemove(t *testing.T) {
	req, err := http.NewRequest("GET", "/remove", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Remove)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "remove", rr.Body.String())
}
