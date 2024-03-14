package views

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

// TODO: Move Secret to ENV
// https://github.com/hay-i/chronologger/issues/35
var SessionStore = sessions.NewCookieStore([]byte("secret"))

func GetFlash(c echo.Context) []interface{} {
	sess, _ := SessionStore.Get(c.Request(), "session")

	return sess.Flashes()
}

func SaveFlash(c echo.Context, flash string) {
	session, _ := SessionStore.Get(c.Request(), "session")

	if flashes := session.Flashes(); len(flashes) > 0 {
		// TODO: Delete old flashes?
	} else {
		session.AddFlash(flash)
	}

	session.Save(c.Request(), c.Response())
}
