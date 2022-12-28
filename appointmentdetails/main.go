package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	appointmentdetails "github.com/pluralsight/webservices/appointmentdetails/appointmentdetailsdata"
	"github.com/pluralsight/webservices/database"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	appointmentdetails.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5004", nil))
}
