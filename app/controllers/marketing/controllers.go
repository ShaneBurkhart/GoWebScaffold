package marketing

import (
	"github.com/ShaneBurkhart/PlanSource/templates"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "marketing/index", nil)
}
