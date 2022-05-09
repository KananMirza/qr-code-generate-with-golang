package handlers

import (
	"html/template"
	"net/http"
	"qr-code-generate-with-golang/pkg/config"
	"qr-code-generate-with-golang/pkg/helpers"
)

func QrCodeHandlers(w http.ResponseWriter, r *http.Request) {
	data := config.GetQrCodeSession(r)
	t, err := template.ParseFiles("views/qr-code.html")
	helpers.CheckError(err)
	helpers.CheckError(t.Execute(w, data))
}
