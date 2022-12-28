package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pluralsight/webservices/database"
	doctor "github.com/pluralsight/webservices/doctor/doctordata"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	doctor.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5002", nil))
}
