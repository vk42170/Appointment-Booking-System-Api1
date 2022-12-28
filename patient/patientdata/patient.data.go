package patient

import (
	"context"
	//"errors"
	"database/sql"
	"log"
	"time"

	"github.com/pluralsight/webservices/database"
)

func getPatientList() ([]Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT
	doctorId,
	patientId,
	patientname, 
	mobile,
	email,
	address 
	FROM patient`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	patients := make([]Patient, 0)
	for results.Next() {
		var patient Patient
		results.Scan(
			&patient.DoctorID,
			&patient.PatientID,
			&patient.Patientname,
			&patient.Mobile,
			&patient.Email,
			&patient.Address)

		patients = append(patients, patient)
	}
	return patients, nil
}

func insertAppointment(patient Patient) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO patient  
	(doctorId,
		patientname, 
		mobile,
		email,
		address) VALUES (?, ?, ?, ?, ?)`,
		patient.DoctorID,
		patient.Patientname,
		patient.Mobile,
		patient.Email,
		patient.Address)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(insertID), nil
}

func registerPatient(patientReg PatientReg) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO pregister  
		(email,
		pass,
		patientname,
		mobile,
		address) VALUES (?, ?, ?, ?, ?)`,
		patientReg.Email,
		patientReg.Pass,
		patientReg.Patientname,
		patientReg.Mobile,
		patientReg.Address)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	return int(insertID), nil
}

func getpatientLogin(patientID int) (*PatientReg, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT
	email,
	pass
	FROM pregister 
	WHERE patientId = ?`, patientID)

	patientlog := &PatientReg{}
	err := row.Scan(
		&patientlog.Email,
		&patientlog.Pass,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return patientlog, nil
}
