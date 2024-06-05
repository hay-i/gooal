package views

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type Flashes map[FlashType][]string

type FlashType string

const (
	FlashError   FlashType = "error"
	FlashWarning FlashType = "warning"
	FlashInfo    FlashType = "info"
	FlashSuccess FlashType = "success"
)

// TODO: Move Secret to ENV
// https://github.com/hay-i/gooal/issues/35
var SessionStore = sessions.NewCookieStore([]byte("secret"))

func getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := SessionStore.Get(r, "my-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic("No session!")
	}

	return session
}

func GetFlashes(c echo.Context) Flashes {
	session := getSession(c.Response(), c.Request())

	defer session.Save(c.Request(), c.Response())

	flashMap := make(Flashes)

	for _, flashType := range []FlashType{FlashError, FlashWarning, FlashInfo, FlashSuccess} {
		flashMessages := []string{}

		for _, flash := range session.Flashes(string(flashType)) {
			flashMessages = append(flashMessages, flash.(string))
		}

		flashMap[flashType] = flashMessages
	}

	return flashMap
}

func AddFlash(c echo.Context, message string, flashType FlashType) {
	session := getSession(c.Response(), c.Request())

	session.AddFlash(message, string(flashType))

	session.Save(c.Request(), c.Response())
}
