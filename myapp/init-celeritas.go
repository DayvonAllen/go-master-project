package main

import (
	celeritas "example"
	"log"
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

	cel.InfoLog.Println("Debug is set to", cel.Debug)

	app := &application{
		App: cel,
	}

	return app
}
