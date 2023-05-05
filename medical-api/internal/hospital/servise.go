package hospital

import (
	"fmt"
	"medical-api/internal/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

type ExistError struct{}

func (*ExistError) Error() string {
	return "This hospital is already registered!"
}

func presentedHospitalDataFromRow(row rowScaner) (*PresentedHospitalData, error) {
	data := &PresentedHospitalData{}
	err := row.Scan(&data.Id, &data.Name, &data.City, &data.Address, &data.Phone, &data.Email, &data.Lon, &data.Lat)
	return data, err
}

type Service interface {
	AddHospital(data MainHospitalData) (int64, error)
	GetHospitalById(id int64) (*PresentedHospitalData, error)
	GetAllHospitals(whreCase string, limitOfset ...int) (*[]PresentedHospitalData, error)
}

type service struct{}

func (s *service) CheckHospitalExist(data MainHospitalData) bool {
	whereCase := fmt.Sprintf(`name = "%s" OR (lon = %g AND lat = %g AND (lon <> 0 AND lat <> 0))`, data.Name, data.Lon, data.Lat)
	result, _ := s.GetAllHospitals(whereCase, 1)
	return len(*result) != 0
}

func (s *service) AddHospital(data MainHospitalData) (int64, error) {
	db := database.Get_db()
	defer db.Close()
	if s.CheckHospitalExist(data) {
		return -1, &ExistError{}
	}
	stmt := fmt.Sprintf("INSERT INTO hospital (name, city, address, phone, email, lon, lat) VALUES (%s)", data.ToString())

	result, err := db.Exec(stmt)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (*service) GetHospitalById(id int64) (*PresentedHospitalData, error) {
	db := database.Get_db()
	defer db.Close()

	stmt := fmt.Sprintf("SELECT * FROM hospital WHERE id=%d", id)
	row := db.QueryRow(stmt)
	return presentedHospitalDataFromRow(row)
}

func (*service) GetAllHospitals(whreCase string, limitOfset ...int) (*[]PresentedHospitalData, error) {
	db := database.Get_db()
	defer db.Close()

	stmt := "SELECT * FROM hospital"
	if whreCase != "" {
		stmt += " WHERE " + whreCase
	}
	if len(limitOfset) >= 1 {
		stmt += fmt.Sprintf(" LIMIT %d", limitOfset[0])
	}
	if len(limitOfset) >= 2 {
		stmt += fmt.Sprintf(", %d", limitOfset[1])
	}

	hospitalsData := []PresentedHospitalData{}
	rows, err := db.Query(stmt)
	defer rows.Close()

	if err != nil {
		return &hospitalsData, err
	}
	for rows.Next() {
		hospitalData, err := presentedHospitalDataFromRow(rows)
		if err != nil {
			return &hospitalsData, err
		}
		hospitalsData = append(hospitalsData, *hospitalData)
	}
	fmt.Println(hospitalsData)
	return &hospitalsData, nil
}

func GetServise() Service {
	return &service{}
}
