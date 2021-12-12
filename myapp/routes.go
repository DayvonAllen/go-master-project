package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"myapp/data"
	"net/http"
	"strconv"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions", a.Handlers.SessionTest)

	a.App.Routes.Get("/create-user", func(writer http.ResponseWriter, request *http.Request) {
		u := data.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "jdoe@gmail.com",
			Active:    1,
			Password:  "password",
		}

		id, err := a.Models.Users.Insert(&u)

		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		_, err = fmt.Fprintf(writer, "%d: %s", id, u.FirstName)
		if err != nil {
			return
		}
	})

	a.App.Routes.Get("/get-all-users", func(writer http.ResponseWriter, request *http.Request) {
		users, err := a.Models.Users.GetAll()

		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		for _, x := range users {
			_, err := fmt.Fprint(writer, x.LastName)
			if err != nil {
				return
			}
		}
	})

	a.App.Routes.Get("/get-user/{id}", func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(request, "id"))

		user, err := a.Models.Users.GetById(id)

		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		_, err = fmt.Fprint(writer, user.LastName)
		if err != nil {
			return
		}
	})

	a.App.Routes.Get("/update-user/{id}", func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(request, "id"))

		user, err := a.Models.Users.GetById(id)

		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		user.LastName = "Johnson"

		err = user.Update(user)

		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}

		_, err = fmt.Fprint(writer, user.LastName)
		if err != nil {
			return
		}
	})

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
