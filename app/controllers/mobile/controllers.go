package mobile

import (
	"github.com/ShaneBurkhart/PlanSource/templates"
	"net/http"
)

func MobileController(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "mobile/index", nil)
}
