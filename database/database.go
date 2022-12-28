package database

import (
	"database/sql"
	"log"
	"time"
)

var DbConn *sql.DB

// SetupDatabase
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:Jaimatadi@123@/appointmentdb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(8)
	DbConn.SetMaxIdleConns(8)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
