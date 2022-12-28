package bookingappointment

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/pluralsight/webservices/cors"
)

const appointmentbookingsPath = "bookingappointment"

func HandleBookingAppointment(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", appointmentbookingsPath))
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
	switch r.Method {
	case http.MethodGet:
		product, err := getBookedAppointment(doctorID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if product == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPut:
		var bookingappointment BookingAppointment
		err := json.NewDecoder(r.Body).Decode(&bookingappointment)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if *bookingappointment.PatientID != doctorID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	appointmentbookingHandler := http.HandlerFunc(HandleBookingAppointment)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, appointmentbookingsPath), cors.Middleware(appointmentbookingHandler))
}
