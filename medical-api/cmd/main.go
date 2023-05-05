package main

import (
	//"medical-api/internal/database"

	"medical-api/internal/app"
)

func main() {
	app := app.NewApp()
	app.Engine(":8000")
}
