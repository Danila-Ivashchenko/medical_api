package testscases

import (
	"fmt"
	"medical-api/internal/user"
)

func TestUsers() {
	userData := user.MainUserData{
		Name:       "Данила",
		Surname:    "Иващенко",
		Patronymic: "Максимович",
		Polis:      "1234567890123456",
		Phone:      "8983313133",
		Email:      "danil2003iva@mail.ru",
		City:       "Иркутск",
		Address:    "Микрорайон Ершовский, дом 180, квартира 31",
		Birthday:   "2003-02-11",
	}
	s := user.GetServise()
	id, err := s.AddUser(userData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)

	id, err = s.AddUser(userData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)

	userData1 := user.MainUserData{
		Name:       "Фёдор",
		Surname:    "Кузнецов",
		Patronymic: "Денисович",
		Polis:      "1234567890123456",
		Phone:      "8999313133",
		Email:      "fedmog@mail.ru",
		City:       "Иркутск",
		Address:    "Микрорайон Приморский, дом 180, квартира 31",
		Birthday:   "2003-09-10",
	}
	id, err = s.AddUser(userData1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)

	userData2 := user.MainUserData{
		Name:       "Фёдор",
		Surname:    "Кузнецов",
		Patronymic: "Денисович",
		Polis:      "1234567890123457",
		Phone:      "8999313133",
		Email:      "fedmog@mail.ru",
		City:       "Иркутск",
		Address:    "Микрорайон Приморский, дом 180, квартира 31",
		Birthday:   "2003-09-10",
	}
	id, err = s.AddUser(userData2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
}
