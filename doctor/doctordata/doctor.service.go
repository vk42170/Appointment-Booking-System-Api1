package doctor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/pluralsight/webservices/cors"
)

const doctorsPath = "doctors"
const doctorregPath = "dregister"
const doctorloginPath = "doctorlogin"

func HandleDoctors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doctorList, err := getDoctorList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(doctorList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var doctorReg DoctorReg
		err := json.NewDecoder(r.Body).Decode(&doctorReg)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		doctorID, err := registerDoctor(doctorReg)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"doctorId":%d}`, doctorID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleDoctorLogin(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", doctorloginPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	doctorID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println(doctorID)
	switch r.Method {
	case http.MethodGet:
		doctor, err := getDoctorLogin(doctorID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if doctor == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(doctor)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	doctorregHandler := http.HandlerFunc(HandleRegister)
	doctorsHandler := http.HandlerFunc(HandleDoctors)
	doctorloginHandler := http.HandlerFunc(HandleDoctorLogin)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, doctorsPath), cors.Middleware(doctorsHandler))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, doctorregPath), cors.Middleware(doctorregHandler))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, doctorloginPath), cors.Middleware(doctorloginHandler))

}
