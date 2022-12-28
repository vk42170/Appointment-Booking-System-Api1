package patient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlePatients(t *testing.T) { //correct
	rec := httptest.NewRecorder()
	PostBody := map[string]interface{}{
		"patientId":   1,
		"doctorId":    1,
		"patientname": "Nirmala Thakur",
		"mobile":      "9966666990",
		"email":       "Nirmala@Thakur.com",
		"address":     "Patna",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5005/api/patients", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	HandlePatients(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}

func TestHandleAppointments(t *testing.T) { //correct
	rec := httptest.NewRecorder()
	res := rec.Result()
	PostBody := map[string]interface{}{
		"doctorId":    5,
		"patientname": "Muruga Prasad",
		"mobile":      "9934987662",
		"email":       "Muruga@Prasad.com",
		"address":     "Siwan mustaff nager",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:5005/api/patientsAppointment", bytes.NewReader(body))
	if err == nil {
		defer res.Body.Close()
	} else {
		t.Fatal(err)
	}

	HandleAppointments(rec, req)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}

func TestHandleRegister(t *testing.T) { //correct
	rec := httptest.NewRecorder()
	res := rec.Result()
	PostBody := map[string]interface{}{
		"email":       "Muruga@Prasad.com",
		"pass":        "Muruga@123",
		"patientname": "Muruga Prasad",
		"mobile":      "9934987662",
		"address":     "Siwan mustaff nager",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodGet, "http://localhost:5005/api/patientRegistration", bytes.NewReader(body))
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

/*
func TestHandlePatientLogin(t *testing.T) { //err

	PostBody := map[string]interface{}{
		"username": "Vishal Kumar",
		"password": "Vishal123",
	}
	body, _ := json.Marshal(PostBody)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:5005/api/patientlogin", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	HandlePatientLogin(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
	}
}
*/
