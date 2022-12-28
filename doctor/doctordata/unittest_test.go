package doctor

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleDoctors(t *testing.T) { //correct
	rec := httptest.NewRecorder()
	PostBody := map[string]interface{}{
		"doctorId": 2,
		"date":     "2022-09-09",
		"slot":     "1PM-2PM",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5002/api/doctors", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	HandleDoctors(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}

func TestHandleRegister(t *testing.T) { //correct
	rec := httptest.NewRecorder()
	res := rec.Result()
	PostBody := map[string]interface{}{
		"email": "Chitranjan@Prasad.com",
		"pass":  "Chitranjan@123",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:5002/api/dregister", bytes.NewReader(body))
	if err == nil {
		defer res.Body.Close()
	} else {
		t.Fatal(err)
	}

	HandleRegister(rec, req)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}
