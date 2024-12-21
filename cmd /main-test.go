package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleCalculate(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Valid expression",
			body:       `{"expression": "2+2*2"}`,
			wantStatus: http.StatusOK,
			wantBody:   `{"result":6}`, 
		},
		{
			name:       "Invalid JSON",
			body:       `{"expression": 2+2*2}`,
			wantStatus: http.StatusUnprocessableEntity,
			wantBody:   `{"error":"Invalid JSON format"}`, 
		},
		{
			name:       "Invalid expression",
			body:       `{"expression": "2+2*"}`,
			wantStatus: http.StatusUnprocessableEntity,
			wantBody:   `{"error":"Expression is not valid"}`, 
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/calculate", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			HandleCalculate(w, req)

			res := w.Result()
			body, _ := io.ReadAll(res.Body)

			if res.StatusCode != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, res.StatusCode)
			}

			if strings.TrimSpace(string(body)) != tt.wantBody { 
				t.Errorf("expected body %q, got %q", tt.wantBody, string(body))
			}
		})
	}
}

func HandleCalculate(w *httptest.ResponseRecorder, req *http.Request) {
	panic("unimplemented")
}
