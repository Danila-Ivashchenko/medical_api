package attachment

import (
	"fmt"
)

type MainAttachmentData struct {
	UserId     int64  `json:"user_id"`
	HospitalID int64  `json:"hospital_id"`
	Date       string `json:"date"`
}

func (m *MainAttachmentData) ToString() string {
	return fmt.Sprintf(`%d, %d, "%s"`, m.UserId, m.HospitalID, m.Date)
}

type PresentedAttachmentData struct {
	Id int64 `json:"id"`
	MainAttachmentData
}

func (m *PresentedAttachmentData) ToString() string {
	return fmt.Sprintf(`%d, %d, %d, "%s"`, m.Id, m.UserId, m.HospitalID, m.Date)
}
