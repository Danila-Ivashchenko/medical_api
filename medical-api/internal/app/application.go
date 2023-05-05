package app

import (
	"medical-api/internal/agregator"
	"medical-api/internal/attachment"
	"medical-api/internal/diseases"
	"medical-api/internal/endpoints"
	"medical-api/internal/hospital"
	"medical-api/internal/user"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *echo.Echo
	endp endpoints.Endpoints
	a    agregator.Agregator
	us   user.Service
	hs   hospital.Service
	as   attachment.Service
	ds   diseases.Service
}

func NewApp() App {
	us := user.GetServise()
	hs := hospital.GetServise()
	as := attachment.GetServise()
	ds := diseases.GetServise()
	a := agregator.GetAgregator(&as, &ds, &hs)
	app := App{a: a, us: us, hs: hs, as: as, ds: ds}
	app.endp = *endpoints.NewEndpoints(&app.us, &app.hs, &app.as, &app.ds, &app.a)
	app.e = echo.New()
	app.routing()
	return app
}

func (app *App) routing() {
	// администрирование
	app.e.POST("/admin/user/all", app.endp.GetAllUsers)
	app.e.POST("/admin/hospital/all", app.endp.GetAllHospitals)
	app.e.POST("/admin/attechment/all", app.endp.GetAttachment)
	app.e.POST("/admin/diseases/all", app.endp.GetDiseasess)
	//endpoints for user
	app.e.POST("/user/register", app.endp.RegisterUser)
	app.e.POST("/user", app.endp.GetUserData)
	app.e.POST("/user/diseases", app.endp.GetUsersDeseases)
	app.e.POST("/user/hospital", app.endp.GetUserHospital)
	//endoints for hospital
	app.e.POST("/hospital", app.endp.GetHospitalData)
	//endpoints for disease
	app.e.POST("disease/add", app.endp.AddDisease)
}

func (app *App) Engine(port string) {
	app.e.Logger.Fatal(app.e.Start(port))
}
