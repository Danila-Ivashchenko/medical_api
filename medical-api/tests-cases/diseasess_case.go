package testscases

import (
	"fmt"
	"medical-api/internal/diseases"
	"medical-api/internal/user"
	"time"
)

func TestDiseasess() {
	us := user.GetServise()
	ds := diseases.GetServise()

	users, _ := us.GetAllUsers("")

	diseasesNames := []string{"Гайморит", "Атит"}
	for i := range *users {
		diseaseData := diseases.MainDiseasesData{
			UserId:     (*users)[i].Id,
			Name:       diseasesNames[i],
			IllDate:    "2023-04-23",
			RecordDate: time.Now().Format("2006-01-02"),
			Status:     true,
		}
		id, err := ds.AddDiseases(diseaseData)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(id)
	}
	ds.ChangeStatusById(1)
	diseases, _ := ds.GetAllDiseasess("")
	for _, data := range *diseases {
		fmt.Println(data)
	}
}
