package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	bookingappointment "github.com/pluralsight/webservices/bookingappointment/bookingappointmentdata"
	"github.com/pluralsight/webservices/database"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	bookingappointment.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5003", nil))
}
