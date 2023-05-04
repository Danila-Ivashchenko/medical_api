package testscases

import (
	"fmt"
	"medical-api/internal/attachment"
	"medical-api/internal/hospital"
	"medical-api/internal/user"
	"time"
)

func TestAttachment() {
	us := user.GetServise()
	hs := hospital.GetServise()
	as := attachment.GetServise()

	users, _ := us.GetAllUsers("")
	hospitals, _ := hs.GetAllHospitals("", 1)

	for i := range *users {
		date := time.Now()
		attachmentData := attachment.MainAttachmentData{
			UserId:     (*users)[i].Id,
			HospitalID: (*hospitals)[0].Id,
			Date:       date.Format("2006-01-02"),
		}
		id, err := as.AddAttachment(attachmentData)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(id)
	}
	attachments, _ := as.GetAllAttachments("")
	for _, data := range *attachments {
		fmt.Println(data)
	}
}
