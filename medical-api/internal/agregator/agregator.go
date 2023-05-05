package agregator

import (
	"fmt"
	"medical-api/internal/attachment"
	"medical-api/internal/diseases"
	"medical-api/internal/hospital"
	"medical-api/internal/user"
)

type Agregator interface {
	ShowUsersDiseases(id int64, limitOfset ...int) (*[]diseases.PresentedDiseasesData, error)
	ShowUsersActiveDiseases(id int64, limitOfset ...int) (*[]diseases.PresentedDiseasesData, error)
	GetUserHospital(id int64) (*hospital.PresentedHospitalData, error)
}

type agregator struct {
	Us *user.Service
	Ds *diseases.Service
	Hs  *hospital.Service
	As *attachment.Service
}

// user - diseases

func (a *agregator) ShowUsersDiseases(id int64, limitOfset ...int) (*[]diseases.PresentedDiseasesData, error) {
	whereCase := fmt.Sprintf("user_id=%d", id)
	return (*a.Ds).GetAllDiseasess(whereCase, limitOfset...)
}

func (a *agregator) ShowUsersActiveDiseases(id int64, limitOfset ...int) (*[]diseases.PresentedDiseasesData, error) {
	whereCase := fmt.Sprintf("user_id=%d AND status = true", id)
	return (*a.Ds).GetAllDiseasess(whereCase, limitOfset...)
}

// user - hospital

func (a *agregator) GetUserHospital(id int64) (*hospital.PresentedHospitalData, error) {
	attachmentData, err := (*a.As).GetUserAttachment(id)
	if err != nil {
		return nil, err
	}
	return (*a.Hs).GetHospitalById(attachmentData.HospitalID)
}

func newAgregator(as *attachment.Service, ds *diseases.Service, hs *hospital.Service) *agregator {
	a := &agregator{}
	a.As = as
	a.Ds = ds
	a.Hs = hs
	//a.Us = user.GetServise()
	return a

}

func GetAgregator(as *attachment.Service, ds *diseases.Service, hs *hospital.Service) Agregator {
	return newAgregator(as, ds, hs)
}
