package views

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

// TODO: Move Secret to ENV
// https://github.com/hay-i/chronologger/issues/35
var SessionStore = sessions.NewCookieStore([]byte("secret"))

func getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := SessionStore.Get(r, "my-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic("No session!")
	}

	return session
}

func GetFlashes(c echo.Context) []interface{} {
	session := getSession(c.Response(), c.Request())

	defer session.Save(c.Request(), c.Response())

	return session.Flashes()
}

func AddFlash(c echo.Context, flash string) {
	session := getSession(c.Response(), c.Request())

	session.AddFlash(flash)

	session.Save(c.Request(), c.Response())
}
