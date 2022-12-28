package appointmentdetails

// AppointmentDetails
type AppointmentDetails struct {
	PatientID   *int   `json:"patientId"`
	DoctorID    int    `json:"doctorId"`
	Patientname string `json:"patientname"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Date        string `json:"date"`
	Slot        string `json:"slot"`
}
