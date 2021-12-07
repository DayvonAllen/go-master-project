package main

import (
	celeritas "example"
	"log"
	"myapp/handlers"
	"os"
)

func initApplication() *application {
	// where did this app start
	// get working directory
	path, err := os.Getwd()

	// cannot find working directory
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := new(celeritas.Celeritas)
	err = cel.New(path)

	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"

	myHandlers := &handlers.Handlers{
		App: cel,
	}

	app := &application{
		App:      cel,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
