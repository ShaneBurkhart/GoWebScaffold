package config

import (
	"github.com/ShaneBurkhart/PlanSource/helpers"
	"net/http"
)

func UpdateLastSeen(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	u := helpers.CurrentUser(r)
	if u != nil {
		u.UpdateLastSeen()
	}
	next(w, r)
}
