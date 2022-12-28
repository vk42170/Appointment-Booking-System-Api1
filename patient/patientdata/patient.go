package patient

type Patient struct {
	PatientID   *int   `json:"patientId"`
	DoctorID    int    `json:"doctorId"`
	Patientname string `json:"patientname"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}
