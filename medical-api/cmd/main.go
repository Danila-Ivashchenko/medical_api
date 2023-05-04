package main

import (
	//"medical-api/internal/database"
	"fmt"
	"medical-api/internal/user"
)

func main() {
	// userData := user.MainUserData{
	// 	Name:       "Danila",
	// 	Surname:    "Ivashchenko",
	// 	Patronymic: "Maksimovich",
	// 	Polis:      "1234567890123456",
	// 	Phone:      "89834076134",
	// 	Email:      "danil2003iva@mail.ru",
	// 	City:       "Irkutsk",
	// 	Address:    "Mkr Ershovsky, d 180, k 31",
	// 	Birthday:   "11.02.2003",
	// }
	s := user.GetServise()
	// id, err := s.AddUserByID(userData)
	// if err != nil{
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(id)
	fmt.Println(s.GetUserById(1))
}
