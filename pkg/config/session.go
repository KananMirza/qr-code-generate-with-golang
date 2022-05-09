package config

import (
	"github.com/gorilla/sessions"
	"net/http"
	"qr-code-generate-with-golang/pkg/helpers"
)

var store = sessions.NewCookieStore([]byte("abcdeffedcbaaabb"))

func SetAlertSession(r *http.Request, w http.ResponseWriter, check bool, msg []string) {
	session, err := store.Get(r, "alert")
	helpers.CheckError(err)
	session.Values["check"] = check
	session.Values["message"] = msg
	helpers.CheckError(session.Save(r, w))
}

func GetAlertSession(r *http.Request) (interface{}, interface{}) {
	session, err := store.Get(r, "alert")
	helpers.CheckError(err)
	return session.Values["check"], session.Values["message"]
}

func SetQrCodeSession(r *http.Request, w http.ResponseWriter, check bool, photo string) {
	session, err := store.Get(r, "qrcode")
	helpers.CheckError(err)
	session.Values["check"] = check
	session.Values["photo"] = photo
	helpers.CheckError(session.Save(r, w))
}

func GetQrCodeSession(r *http.Request) map[string]interface{} {
	session, err := store.Get(r, "qrcode")
	helpers.CheckError(err)
	return map[string]interface{}{
		"check": session.Values["check"],
		"photo": session.Values["photo"],
	}
}
