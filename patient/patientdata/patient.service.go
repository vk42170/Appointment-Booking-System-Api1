package patient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/pluralsight/webservices/cors"
)

const loginPath = "patientlogin"
const patientsPath = "patients"
const patientsAppointPath = "patientsAppointment"
const patientsRegisterPath = "patientRegistration"

func HandlePatients(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		patientList, err := getPatientList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(patientList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func HandleAppointments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var patient Patient
		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		patientID, err := insertAppointment(patient)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"patientId":%d}`, patientID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var patientReg PatientReg
		err := json.NewDecoder(r.Body).Decode(&patientReg)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		patientID, err := registerPatient(patientReg)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"patientId":%d}`, patientID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandlePatientLogin(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", loginPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	patientID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(patientID)
	switch r.Method {
	case http.MethodGet:
		patient, err := getpatientLogin(patientID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if patient == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(patient)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal("Patient not exixts!")
		}
	}
}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {

	patientsHandler := http.HandlerFunc(HandlePatients)
	appointmentHandler := http.HandlerFunc(HandleAppointments)
	pregsiterHandler := http.HandlerFunc(HandleRegister)
	patientloginHandler := http.HandlerFunc(HandlePatientLogin)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, patientsPath), cors.Middleware(patientsHandler))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, patientsAppointPath), cors.Middleware(appointmentHandler))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, patientsRegisterPath), cors.Middleware(pregsiterHandler))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, loginPath), cors.Middleware(patientloginHandler))

}
