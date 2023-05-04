package diseases

import (
	"fmt"
)

type MainDiseasesData struct {
	UserId     int64  `json:"user_id"`
	Name       string `json:"name"`
	IllDate    string `json:"ill_date"`
	RecordDate string `json:"record_date"`
	Status     bool   `json:"status"`
}

func (m *MainDiseasesData) ToString() string {
	return fmt.Sprintf(`%d, "%s", "%s", "%s", %t`, m.UserId, m.Name, m.IllDate, m.RecordDate, m.Status)
}

type PresentedDiseasesData struct {
	Id int64 `json:"id"`
	MainDiseasesData
}

func (m *PresentedDiseasesData) ToString() string {
	return fmt.Sprintf(`%d, %d, %s, "%s", "%s", %t`, m.Id, m.UserId, m.Name, m.IllDate, m.RecordDate, m.Status)
}
