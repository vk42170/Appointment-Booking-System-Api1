package appointmentdetails

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleAppointmentDetails(t *testing.T) { //correct
	rec := httptest.NewRecorder()
	PostBody := map[string]interface{}{
		"patientId":   12,
		"doctorId":    3,
		"patientname": "Vikash Kumar2",
		"mobile":      "8079859243",
		"email":       "Vikash@Kumar1.com",
		"address":     "Allahabad",
		"date":        "2022-08-10",
		"slot":        "4:30PM-6:20PM",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5004/api/appointmentdetails/12", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	HandleAppointmentDetails(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}

func TestHandleAppointmentDetail(t *testing.T) { //correct
	rec := httptest.NewRecorder()
	PostBody := map[string]interface{}{
		"patientId":   13,
		"patientname": "Muruga Prasad2",
		"mobile":      "9934987662",
		"email":       "Muruga@Prasad2.com",
		"address":     "Siwan mustaff nager",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5004/api/appointmentdetail", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	HandleAppointmentDetail(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}
