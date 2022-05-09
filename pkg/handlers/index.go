package handlers

import (
	validator "github.com/elvin-tacirzade/golang-validator"
	"html/template"
	"net/http"
	"qr-code-generate-with-golang/pkg/config"
	"qr-code-generate-with-golang/pkg/helpers"
	"qr-code-generate-with-golang/pkg/models"
)

func MainHandlers(w http.ResponseWriter, r *http.Request) {
	data := models.Page{}
	chk, msg := config.GetAlertSession(r)
	if chk == true {
		data.IsAlert = true
		data.AlertClass = "danger"
		data.AlertTitle = "Error"
		data.AlertContent = msg
		config.SetAlertSession(r, w, false, nil)
	}

	t, err := template.ParseFiles("views/index.html")
	helpers.CheckError(err)
	helpers.CheckError(t.Execute(w, data))
}

func MainHandlersPost(w http.ResponseWriter, r *http.Request) {
	root := "/"
	check := true

	data := r.PostFormValue("data")
	validate := map[string][]string{
		"data": {"required", "max:2000"},
	}

	msg := validator.New(r, validate)
	if len(msg) == 0 {
		root = "/qr-code"
		check = false
		src := "https://chart.googleapis.com/chart?cht=qr&chs=500x500&chl=" + data
		config.SetQrCodeSession(r, w, true, src)
	}
	config.SetAlertSession(r, w, check, msg)
	http.Redirect(w, r, root, http.StatusSeeOther)
}
