package admin

import (
	"github.com/ShaneBurkhart/PlanSource/flash"
	"github.com/ShaneBurkhart/PlanSource/models/user"
	"github.com/ShaneBurkhart/PlanSource/templates"
	"github.com/ShaneBurkhart/PlanSource/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func UserIndexController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	c := r.Form.Get("c")
	o := r.Form.Get("o")
	// Next order value
	var n string = "ASC"
	if o != "DESC" {
		// Everything else but DESC evals to ASC so we only need to set DESC
		// if it isn't DESC already.
		n = "DESC"
	}
	templates.Render(w, r, "admin/user/index", map[string]interface{}{
		"c":     c,
		"o":     o,
		"n":     n,
		"Users": user.AllSorted(c, o),
	})
}

func UserEditController(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	u := user.Find(id)
	// TODO Render 404 if user not found.
	templates.Render(w, r, "admin/user/edit", map[string]interface{}{
		"User": u,
	})
}

func UserUpdateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := getId(r)
	u := user.Find(id)
	updateAttributes(u, r)

	if u.Save() {
		flash.Success(r, "Account successfully updated!")
		http.Redirect(w, r, "/admin/users", http.StatusFound)
	} else {
		data := make(map[string]interface{})
		data["User"] = u
		// Render form to fix
		templates.Render(w, r, "admin/user/edit", map[string]interface{}{
			"User": u,
		})
	}
}

func UserDeleteController(w http.ResponseWriter, r *http.Request) {
	id := getId(r)
	u := user.Find(id)
	u.Delete()
	http.Redirect(w, r, "/admin/users", http.StatusFound)
}

func getId(r *http.Request) int {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	// This will return a 500 instead of 404 because we have parameter
	// validation in routes.go that only allows numbers for id.  If there
	// is an error here, it is something different and most likely a 500.
	util.CheckErr(err)
	return id
}

func updateAttributes(u *user.User, r *http.Request) {
	// TODO this needs to be changed.  Not a good place to put this.
	if first_name := r.PostForm.Get("user[first_name]"); first_name != "" {
		u.FirstName = first_name
	}
	if last_name := r.PostForm.Get("user[last_name]"); last_name != "" {
		u.LastName = last_name
	}
	if company := r.PostForm.Get("user[company]"); company != "" {
		u.Company = company
	}
	if email := r.PostForm.Get("user[email]"); email != "" {
		u.Email = email
	}
	if role := r.PostForm.Get("user[role]"); role != "" {
		u.Role = role
	}
}
