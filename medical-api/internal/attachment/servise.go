package attachment

import (
	"fmt"
	"medical-api/internal/database"
)

type rowScaner interface {
	Scan(dest ...any) error
}

type ExistError struct{}

func (*ExistError) Error() string {
	return "This Attachment is already registered!"
}

func presentedAttachmentDataFromRow(row rowScaner) (*PresentedAttachmentData, error) {
	data := &PresentedAttachmentData{}
	err := row.Scan(&data.Id, &data.UserId, &data.HospitalID, &data.Date)
	return data, err
}

type Service interface {
	AddAttachment(data MainAttachmentData) (int, error)
	GetAttachmentById(id int) (*PresentedAttachmentData, error)
	GetAllAttachments(whreCase string, limitOfset ...int) (*[]PresentedAttachmentData, error)
}

type service struct{}

func (s *service) CheckAttachmentExist(data MainAttachmentData) bool {
	whereCase := fmt.Sprintf(`user_id = %d AND hospital_id = %d`, data.UserId, data.HospitalID)
	result, _ := s.GetAllAttachments(whereCase, 1)
	return len(*result) != 0
}

func (s *service) AddAttachment(data MainAttachmentData) (int, error) {
	db := database.Get_db()
	defer db.Close()
	if s.CheckAttachmentExist(data) {
		return -1, &ExistError{}
	}
	stmt := fmt.Sprintf("INSERT INTO attachment (user_id, hospital_id, date) VALUES (%s)", data.ToString())

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

func (*service) GetAttachmentById(id int) (*PresentedAttachmentData, error) {
	db := database.Get_db()
	defer db.Close()

	stmt := fmt.Sprintf("SELECT * FROM attachment WHERE id=%d", id)
	row := db.QueryRow(stmt)
	return presentedAttachmentDataFromRow(row)
}

func (*service) GetAllAttachments(whreCase string, limitOfset ...int) (*[]PresentedAttachmentData, error) {
	db := database.Get_db()
	defer db.Close()

	stmt := "SELECT * FROM attachment"
	if whreCase != "" {
		stmt += " WHERE " + whreCase
	}
	if len(limitOfset) >= 1 {
		stmt += fmt.Sprintf(" LIMIT %d", limitOfset[0])
	}
	if len(limitOfset) >= 2 {
		stmt += fmt.Sprintf(", %d", limitOfset[1])
	}

	AttachmentsData := []PresentedAttachmentData{}
	rows, err := db.Query(stmt)
	defer rows.Close()

	if err != nil {
		return &AttachmentsData, err
	}
	for rows.Next() {
		AttachmentData, err := presentedAttachmentDataFromRow(rows)
		if err != nil {
			return &AttachmentsData, err
		}
		AttachmentsData = append(AttachmentsData, *AttachmentData)
	}
	return &AttachmentsData, nil
}

func GetServise() Service {
	return &service{}
}
