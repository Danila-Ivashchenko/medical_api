package endpoints

import (
	"medical-api/internal/diseases"
	"medical-api/internal/hospital"
	"medical-api/internal/user"
)

// requests

type UserIdRequest struct {
	Id int64 `json:"user_id"`
}

type UserDiseasesRequest struct {
	UserIdRequest
	Limit int `json:"limit"`
	Ofset int `json:"ofset"`
}

type HospitalIdRequest struct {
	Id int64 `json:"hospital_id"`
}

type DiseaseAddRequest struct {
	UserIdRequest
	Name       string `json:"name"`
	IllDate    string `json:"ill_date"`
	RecordDate string `json:"record_date"`
	Status     bool   `json:"status"`
}

// responses

type baseResponse struct {
	Exaption map[string]interface{} `json:"exaption"`
	Succsess bool                   `json:"success"`
}

//

func BadBaseResponse(err error) baseResponse {
	response := baseResponse{}
	response.Exaption = map[string]interface{}{"error": err.Error()}
	response.Succsess = false
	return response
}

func GoodBaseResponse() baseResponse {
	response := baseResponse{}
	response.Succsess = true
	return response
}

type HospitalDataResponse struct {
	baseResponse
	HospitalData *hospital.PresentedHospitalData `json:"hospital_data"`
}

//

func BadHospitalResponse(err error) HospitalDataResponse {
	response := HospitalDataResponse{}
	response.Exaption = map[string]interface{}{"error": err.Error()}
	response.Succsess = false
	return response
}

func GoodHospitalResponse(hospitalData *hospital.PresentedHospitalData) HospitalDataResponse {
	response := HospitalDataResponse{}
	response.Succsess = true
	response.HospitalData = hospitalData
	return response
}

type UserDataRespose struct {
	baseResponse
	UserData *user.PresentedUserData `json:"user_data"`
}

//

func BadUserResponse(err error) UserDataRespose {
	response := UserDataRespose{}
	response.Exaption = map[string]interface{}{"error": err.Error()}
	response.Succsess = false
	return response
}

func GoodUserResponse(userData *user.PresentedUserData) UserDataRespose {
	response := UserDataRespose{}
	response.Succsess = true
	response.UserData = userData
	return response
}

//

type DiseasesResponse struct {
	baseResponse
	Diseases *[]diseases.PresentedDiseasesData `json:"diseases"`
}

func BadDiseasesResponse(err error) DiseasesResponse {
	response := DiseasesResponse{}
	response.Exaption = map[string]interface{}{"error": err.Error()}
	response.Succsess = false
	return response
}

func GoodDiseasesResponse(diseasesData *[]diseases.PresentedDiseasesData) DiseasesResponse {
	response := DiseasesResponse{}
	response.Succsess = true
	response.Diseases = diseasesData
	return response
}

type DiseaseResponse struct {
	baseResponse
	Disease *diseases.PresentedDiseasesData
}

func BadDiseaseResponse(err error) DiseaseResponse {
	response := DiseaseResponse{}
	response.Exaption = map[string]interface{}{"error": err.Error()}
	response.Succsess = false
	return response
}

func GoodDiseaseResponse(data *diseases.PresentedDiseasesData) DiseaseResponse {
	response := DiseaseResponse{}
	response.Succsess = true
	response.Disease = data
	return response
}
