package user

import (
	"database/sql"
	"fmt"
	"medical-api/internal/database"
	"net/smtp"
)

func presentedUserDataFromRow(row *sql.Row) (*PresentedUserData, error){
	data := &PresentedUserData{}
	err := row.Scan(&data.Id, &data.Name, &data.Surname, &data.Patronymic, &data.Polis, &data.Phone, &data.Email, &data.City, &data.Address, &data.Birthday)
	return data, err
}

type Service interface {
	AddUserByID(data MainUserData) (int, error)
	GetUserById(id int) (*PresentedUserData, error)
	GetAllUsers(whreCase string) ([]*PresentedUserData, error)
}

type service struct{}

func (*service) AddUserByID(data MainUserData) (int, error) {
	db := database.Get_db()
	defer db.Close()

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

func (*service) GetUserById(id int) (*PresentedUserData, error){
	db := database.Get_db()
	defer db.Close()

	stmt := fmt.Sprintf("SELECT * FROM users WHERE id=%d", id)
	row := db.QueryRow(stmt)
	return presentedUserDataFromRow(row)
}

func (*service) GetAllUsers(whreCase string, limitOfset... int) ([]*PresentedUserData, error) {
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

	rows, err := db.Query(stmt)
	
}



func GetServise() Service {
	return &service{}
}
