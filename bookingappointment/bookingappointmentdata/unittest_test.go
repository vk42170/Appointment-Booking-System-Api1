package bookingappointment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleBookingAppointment(t *testing.T) { //correct
	rec := httptest.NewRecorder()
	PostBody := map[string]interface{}{
		"patientId":   1,
		"doctorId":    1,
		"patientname": "Nirmala Thakur",
		"mobile":      "9966666990",
		"email":       "Nirmala@Thakur.com",
		"address":     "Patna",
		"date":        "2022-08-07",
		"slot":        "12PM-1PM",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5003/api/bookingappointment/1", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	HandleBookingAppointment(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}
