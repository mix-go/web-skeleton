package globals

import (
	"github.com/go-session/session"
	"github.com/mix-go/xdi"
)

func Session() (manager *session.Manager) {
	if err := xdi.Populate("session", &manager); err != nil {
		panic(err)
	}
	return
}
