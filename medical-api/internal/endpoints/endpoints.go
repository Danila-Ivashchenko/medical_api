package endpoints

import (
	"medical-api/internal/agregator"
	"medical-api/internal/attachment"
	"medical-api/internal/diseases"
	"medical-api/internal/hospital"
	"medical-api/internal/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func readBody(с echo.Context) ([]byte, error) {
	bytes := []byte{}
	_, err := с.Request().Body.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

type Endpoints struct {
	a  *agregator.Agregator
	us *user.Service
	hs *hospital.Service
	as *attachment.Service
	ds *diseases.Service
}

// users endpoints

func (endp *Endpoints) RegisterUser(c echo.Context) error {
	userData := user.MainUserData{}
	c.Bind(&userData)
	id, err := (*endp.us).AddUser(userData)
	if err != nil {
		c.JSON(http.StatusOK, BadUserResponse(err))
	}
	data := user.PresentedUserDataFromMain(id, userData)
	return c.JSON(http.StatusOK, GoodUserResponse(&data))
}

func (endp *Endpoints) GetUserData(c echo.Context) error {
	request := UserIdRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadUserResponse(err))
	}
	data, err := (*endp.us).GetUserById(request.Id)
	if err != nil {
		return c.JSON(http.StatusOK, BadUserResponse(err))
	}
	return c.JSON(http.StatusOK, GoodUserResponse(data))
}

func (endp *Endpoints) GetUsersDeseases(c echo.Context) error {
	request := UserIdRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadDiseasesResponse(err))
	}
	data, err := (*endp.a).ShowUsersDiseases(request.Id)
	if err != nil {
		return c.JSON(http.StatusOK, BadDiseasesResponse(err))
	}
	return c.JSON(http.StatusOK, GoodDiseasesResponse(data))
}

func (endp *Endpoints) GetUserHospital(c echo.Context) error {
	request := UserIdRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadDiseasesResponse(err))
	}
	data, err := (*endp.a).GetUserHospital(request.Id)
	if err != nil {
		return c.JSON(http.StatusOK, BadHospitalResponse(err))
	}
	return c.JSON(http.StatusOK, GoodHospitalResponse(data))
}

// hospital endpoints

func (endp *Endpoints) GetHospitalData(c echo.Context) error {
	request := HospitalIdRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadHospitalResponse(err))
	}
	data, err := (*endp.a).GetUserHospital(request.Id)
	if err != nil {
		return c.JSON(http.StatusOK, BadHospitalResponse(err))
	}
	return c.JSON(http.StatusOK, GoodHospitalResponse(data))
}

// disease endpoints

func (endp *Endpoints) AddDisease(c echo.Context) error {
	request := DiseaseAddRequest{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusOK, BadBaseResponse(err))
	}
	diseaseData := diseases.MainDiseasesData{UserId: request.Id,Name: request.Name, IllDate: request.IllDate, RecordDate: request.RecordDate, Status: request.Status}
	id, err := (*endp.ds).AddDiseases(diseaseData)
	if err != nil {
		return c.JSON(http.StatusOK, BadDiseaseResponse(err))
	}

	responseData, err := (*endp.ds).GetDiseasesById(id)
	if err != nil {
		return c.JSON(http.StatusOK, BadDiseaseResponse(err))
	}
	return c.JSON(http.StatusOK, GoodDiseaseResponse(responseData))
}

// функции для администрирования

func (endp *Endpoints) GetAllUsers(c echo.Context) error {
	data, _ := (*endp.us).GetAllUsers("")
	return c.JSON(http.StatusOK, data)
}

func (endp *Endpoints) GetAllHospitals(c echo.Context) error {
	data, _ := (*endp.hs).GetAllHospitals("")
	return c.JSON(http.StatusOK, data)
}

func (endp *Endpoints) GetAttachment(c echo.Context) error {
	data, _ := (*endp.as).GetAllAttachments("")
	return c.JSON(http.StatusOK, data)
}

func (endp *Endpoints) GetDiseasess(c echo.Context) error {
	data, _ := (*endp.ds).GetAllDiseasess("")
	return c.JSON(http.StatusOK, data)
}

func NewEndpoints(us *user.Service, hs *hospital.Service, as *attachment.Service, ds *diseases.Service, a *agregator.Agregator) *Endpoints {
	endp := &Endpoints{}
	endp.us = us
	endp.hs = hs
	endp.as = as
	endp.ds = ds
	endp.a = a
	return endp
}
