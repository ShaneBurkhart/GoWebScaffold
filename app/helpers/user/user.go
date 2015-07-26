package helpers

import (
	"github.com/ShaneBurkhart/PlanSource/models/user"
	"github.com/goincremental/negroni-sessions"
	"net/http"
)

func CurrentUser(r *http.Request) *user.User {
	if id, ok := sessions.GetSession(r).Get("user_id").(int); ok {
		return user.Find(id)
	} else {
		return nil
	}
}

func SignIn(r *http.Request, u user.User) {
	sessions.GetSession(r).Set("user_id", u.Id)
}

func SignOut(r *http.Request) {
	sessions.GetSession(r).Delete("user_id")
}
