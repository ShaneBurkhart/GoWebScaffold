package flash

import (
	"fmt"
	"github.com/goincremental/negroni-sessions"
	"net/http"
)

const (
	// Class names that will go on flash div
	infoClass    = "info"
	errorClass   = "error"
	successClass = "success"
)

type Flash struct {
	Type    string
	Message string
}

func Info(r *http.Request, m string) {
	sessions.GetSession(r).AddFlash(Flash{infoClass, m})
}

func Error(r *http.Request, m string) {
	sessions.GetSession(r).AddFlash(Flash{errorClass, m})
}

func Success(r *http.Request, m string) {
	sessions.GetSession(r).AddFlash(Flash{successClass, m})
}

func Flashes(r *http.Request) []Flash {
	fs := make([]Flash, 0)
	for _, v := range sessions.GetSession(r).Flashes() {
		if f, ok := v.(Flash); ok {
			fs = append(fs, f)
		} else if s, ok := v.(fmt.Stringer); ok {
			fs = append(fs, Flash{infoClass, s.String()})
		}
	}
	return fs
}
