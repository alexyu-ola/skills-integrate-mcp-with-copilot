package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGetWeather_Success(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/weather?city=Beijing", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var info WeatherInfo
	if err := json.Unmarshal(w.Body.Bytes(), &info); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if info.City != "Beijing" {
		t.Errorf("expected city Beijing, got %s", info.City)
	}
}

func TestGetWeather_CaseInsensitive(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/weather?city=SHANGHAI", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var info WeatherInfo
	if err := json.Unmarshal(w.Body.Bytes(), &info); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}
	if info.City != "Shanghai" {
		t.Errorf("expected city Shanghai, got %s", info.City)
	}
}

func TestGetWeather_MissingCity(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/weather", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestGetWeather_UnknownCity(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/weather?city=UnknownCity", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestNormalizeCity(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"Beijing", "beijing"},
		{"SHANGHAI", "shanghai"},
		{"new york", "new york"},
		{"New York", "new york"},
	}
	for _, tc := range cases {
		result := normalizeCity(tc.input)
		if result != tc.expected {
			t.Errorf("normalizeCity(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}
