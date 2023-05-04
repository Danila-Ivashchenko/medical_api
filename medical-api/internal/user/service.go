package user

import (
	"fmt"
	"medical-api/internal/database"
)

type ExistError struct{}

func (*ExistError) Error() string {
	return "This user is already registered!"
}

type rowScaner interface {
	Scan(dest ...any) error
}

func presentedUserDataFromRow(row rowScaner) (*PresentedUserData, error) {
	data := &PresentedUserData{}
	err := row.Scan(&data.Id, &data.Name, &data.Surname, &data.Patronymic, &data.Polis, &data.Phone, &data.Email, &data.City, &data.Address, &data.Birthday)
	return data, err
}

type Service interface {
	AddUserByID(data MainUserData) (int, error)
	GetUserById(id int) (*PresentedUserData, error)
	GetAllUsers(whreCase string, limitOfset ...int) (*[]PresentedUserData, error)
}

type service struct{}

func (s *service) AddUserByID(data MainUserData) (int, error) {
	db := database.Get_db()
	defer db.Close()
	if s.CheckUserExist(data) {
		return -1, &ExistError{}
	}
	stmt := fmt.Sprintf("INSERT INTO users (name, surname, patronymic, polis, phone, email, city, address, birthday) VALUES (%s)", data.ToString())

	result, err := db.Exec(stmt)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (s *service) CheckUserExist(data MainUserData) bool {
	whereCase := fmt.Sprintf(`polis = "%s" OR phone = "%s" OR email = "%s"`, data.Polis, data.Phone, data.Email)
	result, _ := s.GetAllUsers(whereCase, 1)
	return len(*result) != 0
}

func (*service) GetUserById(id int) (*PresentedUserData, error) {
	db := database.Get_db()
	defer db.Close()

	stmt := fmt.Sprintf("SELECT * FROM users WHERE id=%d", id)
	row := db.QueryRow(stmt)
	return presentedUserDataFromRow(row)
}

func (*service) GetAllUsers(whreCase string, limitOfset ...int) (*[]PresentedUserData, error) {
	db := database.Get_db()
	defer db.Close()

	stmt := "SELECT * FROM users"
	if whreCase != "" {
		stmt += " WHERE " + whreCase
	}
	if len(limitOfset) >= 1 {
		stmt += fmt.Sprintf(" LIMIT %d", limitOfset[0])
	}
	if len(limitOfset) >= 2 {
		stmt += fmt.Sprintf(", %d", limitOfset[1])
	}

	usersData := []PresentedUserData{}
	rows, err := db.Query(stmt)
	defer rows.Close()

	if err != nil {
		return &usersData, err
	}
	for rows.Next() {
		userData, err := presentedUserDataFromRow(rows)
		if err != nil {
			return &usersData, err
		}
		usersData = append(usersData, *userData)
	}
	return &usersData, nil
}

func GetServise() Service {
	return &service{}
}
