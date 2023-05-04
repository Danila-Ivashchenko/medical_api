package diseases

import (
	"fmt"
	"medical-api/internal/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

type ExistError struct{}

func (*ExistError) Error() string {
	return "This diseases is already registered!"
}

func presentedDiseasesDataFromRow(row rowScaner) (*PresentedDiseasesData, error) {
	data := &PresentedDiseasesData{}
	err := row.Scan(&data.Id, &data.UserId, &data.Name, &data.IllDate, &data.RecordDate, &data.Status)
	return data, err
}

type Service interface {
	AddDiseases(data MainDiseasesData) (int, error)
	GetDiseasesById(id int) (*PresentedDiseasesData, error)
	GetAllDiseasess(whreCase string, limitOfset ...int) (*[]PresentedDiseasesData, error)
	ChangeStatusById(id int) (*PresentedDiseasesData, error)
}

type service struct{}

func (s *service) CheckDiseasesExist(data MainDiseasesData) bool {
	whereCase := fmt.Sprintf(`user_id = %d AND name = "%s"`, data.UserId, data.Name)
	result, _ := s.GetAllDiseasess(whereCase, 1)
	return len(*result) != 0
}

func (s *service) AddDiseases(data MainDiseasesData) (int, error) {
	db := database.Get_db()
	defer db.Close()
	if s.CheckDiseasesExist(data) {
		return -1, &ExistError{}
	}
	stmt := fmt.Sprintf("INSERT INTO diseases (user_id, name, ill_date, record_date, status) VALUES (%s)", data.ToString())

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

func (*service) GetDiseasesById(id int) (*PresentedDiseasesData, error) {
	db := database.Get_db()
	defer db.Close()

	stmt := fmt.Sprintf("SELECT * FROM diseases WHERE id=%d", id)
	row := db.QueryRow(stmt)
	return presentedDiseasesDataFromRow(row)
}

func (*service) GetAllDiseasess(whreCase string, limitOfset ...int) (*[]PresentedDiseasesData, error) {
	db := database.Get_db()
	defer db.Close()

	stmt := "SELECT * FROM diseases"
	if whreCase != "" {
		stmt += " WHERE " + whreCase
	}
	if len(limitOfset) >= 1 {
		stmt += fmt.Sprintf(" LIMIT %d", limitOfset[0])
	}
	if len(limitOfset) >= 2 {
		stmt += fmt.Sprintf(", %d", limitOfset[1])
	}

	DiseasessData := []PresentedDiseasesData{}
	rows, err := db.Query(stmt)
	defer rows.Close()

	if err != nil {
		return &DiseasessData, err
	}
	for rows.Next() {
		DiseasesData, err := presentedDiseasesDataFromRow(rows)
		if err != nil {
			return &DiseasessData, err
		}
		DiseasessData = append(DiseasessData, *DiseasesData)
	}
	return &DiseasessData, nil
}

func (s *service) ChangeStatusById(id int) (*PresentedDiseasesData, error) {
	_, err := s.GetDiseasesById(id)
	if err != nil {
		return nil, err
	}
	db := database.Get_db()
	defer db.Close()

	stmt := fmt.Sprintf("UPDATE diseases SET status = !status WHERE id = %d", id)
	_, err = db.Exec(stmt)
	if err != nil {
		return nil, err
	}
	return s.GetDiseasesById(id)
}

func GetServise() Service {
	return &service{}
}
