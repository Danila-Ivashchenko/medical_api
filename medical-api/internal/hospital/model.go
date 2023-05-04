package hospital

import (
	"fmt"
)

type MainHospitalData struct {
	Name    string  `json:"name"`
	City    string  `json:"city"`
	Address string  `json:"address"`
	Phone   string  `json:"contact_phone"`
	Email   string  `json:"cantact_email"`
	Lon     float64 `json:"lon"`
	Lat     float64 `json:"lat"`
}

func (m *MainHospitalData) ToString() string {
	return fmt.Sprintf("'%s', '%s', '%s', '%s', '%s', '%g', '%g'", m.Name, m.City, m.Address, m.Phone, m.Email, m.Lon, m.Lat)
}

type PresentedHospitalData struct {
	Id int64 `json:"id"`
	MainHospitalData
}

func (m *PresentedHospitalData) ToString() string {
	return fmt.Sprintf("'%d', '%s', '%s', '%s', '%s', '%s', '%g', '%g'", m.Id, m.Name, m.City, m.Address, m.Phone, m.Email, m.Lon, m.Lat)
}
