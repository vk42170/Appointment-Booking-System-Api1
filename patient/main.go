package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pluralsight/webservices/database"
	patient "github.com/pluralsight/webservices/patient/patientdata"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	patient.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5005", nil))
}
