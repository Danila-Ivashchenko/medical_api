package testscases

import (
	"fmt"
	"medical-api/internal/hospital"
)

func TestHospital() {
	hospitalData1 := hospital.MainHospitalData{
		Name:    "11 поликлиника",
		City:    "Иркутск",
		Address: "Лермонтова 89",
		Phone:   "83952411966",
		Email:   "info-irgp11@mail.ru",
	}
	hospitalData2 := hospital.MainHospitalData{
		Name:    "11 поликлиника",
		City:    "Иркутск",
		Address: "Лермонтова 89",
		Phone:   "83952411966",
		Email:   "info-irgp11@mail.ru",
	}
	hospitalData3 := hospital.MainHospitalData{
		Name:    "12 поликлиника",
		City:    "Иркутск",
		Address: "Лермонтова 12",
		Phone:   "83952411977",
		Email:   "info-irgp12@mail.ru",
	}
	s := hospital.GetServise()
	id, err := s.AddHospital(hospitalData1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
	id, err = s.AddHospital(hospitalData2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
	id, err = s.AddHospital(hospitalData3)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)

	hospitals, err := s.GetAllHospitals("")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, data := range *hospitals {
		fmt.Println(data)
	}
}
