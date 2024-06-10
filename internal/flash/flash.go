package flash

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type Flashes map[Type][]string

type Type string

const (
	Error   Type = "error"
	Warning Type = "warning"
	Info    Type = "info"
	Success Type = "success"
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

func Get(c echo.Context) Flashes {
	session := getSession(c.Response(), c.Request())

	defer session.Save(c.Request(), c.Response())

	flashMap := make(Flashes)

	for _, flashType := range []Type{Error, Warning, Info, Success} {
		flashMessages := []string{}

		for _, flash := range session.Flashes(string(flashType)) {
			flashMessages = append(flashMessages, flash.(string))
		}

		flashMap[flashType] = flashMessages
	}

	return flashMap
}

func Add(c echo.Context, message string, flashType Type) {
	session := getSession(c.Response(), c.Request())

	session.AddFlash(message, string(flashType))

	session.Save(c.Request(), c.Response())
}
