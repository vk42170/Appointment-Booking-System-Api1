package doctor

//Doctor
type Doctor struct {
	DoctorID *int   `json:"doctorId"`
	Date     string `json:"date"`
	Slot     string `json:"slot"`
}
