package appointmentdetails

import (
	"context"
	"database/sql"

	"errors"
	"log"
	"time"

	"github.com/pluralsight/webservices/database"
)

func getAppointmentDetails(patientID int) (*AppointmentDetails, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT
	p.patientId,
	d.doctorId, 
	p.patientname,
	p.mobile,
	p.email, 
	p.address,
	d.date, 
	d.slot 
	FROM doctors d, patient p 
	WHERE d.doctorId = p.doctorId and p.patientId = ?`, patientID)

	appointmentdetails := &AppointmentDetails{}
	err := row.Scan(
		&appointmentdetails.PatientID,
		&appointmentdetails.DoctorID,
		&appointmentdetails.Patientname,
		&appointmentdetails.Mobile,
		&appointmentdetails.Email,
		&appointmentdetails.Address,
		&appointmentdetails.Date,
		&appointmentdetails.Slot,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return appointmentdetails, nil
}

func updateAppointmentDetails(appointmentdetails AppointmentDetails) error {
	// if the patient id is set, update, otherwise add
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if appointmentdetails.PatientID == nil || *appointmentdetails.PatientID == 0 {
		return errors.New("patient has invalid ID")
	}
	_, err := database.DbConn.ExecContext(ctx, `UPDATE patient SET 
		patientname=?, 
		mobile=?, 
		email=?, 
		address=?
		WHERE patientId=?`,
		appointmentdetails.Patientname,
		appointmentdetails.Mobile,
		appointmentdetails.Email,
		appointmentdetails.Address,
		appointmentdetails.PatientID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
