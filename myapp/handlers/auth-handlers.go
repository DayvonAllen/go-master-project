package handlers

import "net/http"

func (h *Handlers) UserLogin(writer http.ResponseWriter, request *http.Request) {
	err := h.App.Render.Page(writer, request, "login", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println(err)
		return
	}
}

func (h *Handlers) PostUserLogin(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()

	if err != nil {
		_, err2 := writer.Write([]byte(err.Error()))
		if err2 != nil {
			h.App.ErrorLog.Println(err2)
		}
		return
	}

	email := request.Form.Get("email")
	password := request.Form.Get("password")

	user, err := h.Models.Users.GetByEmail(email)

	if err != nil {
		_, err2 := writer.Write([]byte(err.Error()))
		if err2 != nil {
			h.App.ErrorLog.Println(err2)
		}
		return
	}
	matches, err := user.PasswordMatches(password)

	if err != nil {
		_, err2 := writer.Write([]byte(err.Error()))
		if err2 != nil {
			h.App.ErrorLog.Println(err2)
			return
		}
		return
	}

	if !matches {
		_, err = writer.Write([]byte("Invalid Password"))
		if err != nil {
			h.App.ErrorLog.Println(err)
		}
		return
	}

	h.App.Session.Put(request.Context(), "userID", user.ID)

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}
