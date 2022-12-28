package doctor

import (
	"context"
	//"errors"
	"database/sql"
	"log"
	"time"

	"github.com/pluralsight/webservices/database"
)

func getDoctorList() ([]Doctor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT
	doctorId, 
	date, 
	slot 
	FROM doctors`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	doctors := make([]Doctor, 0)
	for results.Next() {
		var doctor Doctor
		results.Scan(
			&doctor.DoctorID,
			&doctor.Date,
			&doctor.Slot)

		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

func registerDoctor(doctorReg DoctorReg) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO dregister  
		(email,
		pass) VALUES (?, ?)`,
		doctorReg.Email,
		doctorReg.Pass)
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

func getDoctorLogin(doctorID int) (*Login, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT
	email,
	pass
	FROM dregister 
	WHERE doctorId = ?`, doctorID)

	doctorlog := &Login{}
	err := row.Scan(
		&doctorlog.Email,
		&doctorlog.Pass,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return doctorlog, nil
}
