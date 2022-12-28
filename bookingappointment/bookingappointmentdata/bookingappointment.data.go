package bookingappointment

import (
	"context"
	//"database/sql"

	"log"
	"time"

	"github.com/pluralsight/webservices/database"
)

/*
func getBookedAppointment(doctorID int) (*BookingAppointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT
	d.doctorId,
	p.patientname,
	p.mobile,
	p.email,
	p.address,
	d.date,
	d.slot
	FROM doctors d, patient p
	WHERE d.doctorId = p.doctorId and d.doctorId = ?`, doctorID)

	bookingappointment := &BookingAppointment{}
	err := row.Scan(
		&bookingappointment.DoctorID,
		&bookingappointment.Patientname,
		&bookingappointment.Mobile,
		&bookingappointment.Email,
		&bookingappointment.Address,
		&bookingappointment.Date,
		&bookingappointment.Slot,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return bookingappointment, nil
}*/

func getBookedAppointment(doctorID int) ([]BookingAppointment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT
	p.patientId,
	d.doctorId, 
	p.patientname,
	p.mobile,
	p.email, 
	p.address,
	d.date, 
	d.slot 
	FROM doctors d, patient p 
	WHERE d.doctorId = p.doctorId and d.doctorId = ?`, doctorID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	bookingappointments := make([]BookingAppointment, 0)
	for results.Next() {
		var bookingappointment BookingAppointment
		results.Scan(
			&bookingappointment.PatientID,
			&bookingappointment.DoctorID,
			&bookingappointment.Patientname,
			&bookingappointment.Mobile,
			&bookingappointment.Email,
			&bookingappointment.Address,
			&bookingappointment.Date,
			&bookingappointment.Slot)

		bookingappointments = append(bookingappointments, bookingappointment)
	}
	return bookingappointments, nil
}
