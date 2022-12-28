package appointmentdetails

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/pluralsight/webservices/cors"
)

const appointmentdetailsPath = "appointmentdetails"
const appointmentdetailPath = "appointmentdetail"

func HandleAppointmentDetails(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", appointmentdetailsPath))
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
		doctor, err := getAppointmentDetails(doctorID)
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

func HandleAppointmentDetail(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", appointmentdetailPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodPut:
		var appointmentdetails AppointmentDetails
		err := json.NewDecoder(r.Body).Decode(&appointmentdetails)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateAppointmentDetails(appointmentdetails)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	appointmentbookingHandler := http.HandlerFunc(HandleAppointmentDetails)
	appointmentdetailHandler := http.HandlerFunc(HandleAppointmentDetail)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, appointmentdetailsPath), cors.Middleware(appointmentbookingHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, appointmentdetailPath), cors.Middleware(appointmentdetailHandler))
}
