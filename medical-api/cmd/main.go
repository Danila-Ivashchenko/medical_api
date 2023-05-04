package main

import (
	//"medical-api/internal/database"

	testscases "medical-api/tests-cases"
)

func main() {
	// userData := user.MainUserData{
	// 	Name:       "Фёдор",
	// 	Surname:    "Кузнецов",
	// 	Patronymic: "Денисович",
	// 	Polis:      "1234567890123456",
	// 	Phone:      "8983313133",
	// 	Email:      "fedmog@mail.ru",
	// 	City:       "Иркутск",
	// 	Address:    "Микрорайон Приморский, дом 41, квартира 21",
	// 	Birthday:   "10.09.2003",
	// }
	// s := user.GetServise()
	// id, err := s.AddUserByID(userData)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(id)
	// //fmt.Println(s.GetAllUsers("1", 1))
	//testscases.TestHospital()
	//testscases.TestAttachment()
	testscases.TestDiseasess()
}
