package globals

import (
    "github.com/go-session/session"
    "github.com/mix-go/console"
)

func Session() *session.Manager {
    return console.App.Get("session").(*session.Manager)
}
