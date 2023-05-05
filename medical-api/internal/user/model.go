package user

import "fmt"

type MainUserData struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Polis      string `json:"polis"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	City       string `json:"city"`
	Address    string `json:"address"`
	Birthday   string `json:"birthday"`
}

func (m *MainUserData) ToString() string {
	return fmt.Sprintf("'%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'", m.Name, m.Surname, m.Patronymic, m.Polis, m.Phone, m.Email, m.City, m.Address, m.Birthday)
}

type PresentedUserData struct {
	Id int64 `json:"id"`
	MainUserData
}

func (m *PresentedUserData) ToString() string {
	return fmt.Sprintf("'%d','%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'", m.Id, m.Name, m.Surname, m.Patronymic, m.Polis, m.Phone, m.Email, m.City, m.Address, m.Birthday)
}

func PresentedUserDataFromMain(id int64, data MainUserData) PresentedUserData {
	pdata := PresentedUserData{}
	pdata.Id = id
	pdata.Name = data.Name
	pdata.Surname = data.Surname
	pdata.Patronymic = data.Patronymic
	pdata.Polis = data.Polis
	pdata.Phone = data.Phone
	pdata.Email = data.Email
	pdata.City = data.City
	pdata.Address = data.Address
	pdata.Birthday = data.Birthday
	return pdata
}
