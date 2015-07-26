package user

import (
	"github.com/ShaneBurkhart/PlanSource/flash"
	"github.com/ShaneBurkhart/PlanSource/helpers"
	"github.com/ShaneBurkhart/PlanSource/models/user"
	"github.com/ShaneBurkhart/PlanSource/templates"
	"net/http"
)

func EditController(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["User"] = helpers.CurrentUser(r)
	templates.Render(w, r, "user/edit", data)
}

func UpdateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := helpers.CurrentUser(r)
	updateAttributes(u, r)
	old_password := r.PostForm.Get("user[password]")
	new_password := r.PostForm.Get("user[new_password]")
	new_password_conf := r.PostForm.Get("user[new_password_confirmation]")
	u.UpdatePassword(old_password, new_password, new_password_conf)
	if u.Save() {
		flash.Success(r, "Account successfully updated!")
		http.Redirect(w, r, "/app", http.StatusFound)
	} else {
		// Erase password for good practice.
		u.Password = ""
		data := make(map[string]interface{})
		data["User"] = u
		// Render form to fix
		templates.Render(w, r, "user/edit", data)
	}
}

func CreateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := user.User{
		FirstName: r.PostForm.Get("user[first_name]"),
		LastName:  r.PostForm.Get("user[last_name]"),
		Company:   r.PostForm.Get("user[company]"),
		Email:     r.PostForm.Get("user[email]"),
		Password:  r.PostForm.Get("user[password]"),
		Role:      "viewer",
	}
	if u.Save() {
		helpers.SignIn(r, u)
		flash.Success(r, "Account successfully created!")
		http.Redirect(w, r, "/app", http.StatusFound)
	} else {
		// Erase password for good practice.
		u.Password = ""
		data := make(map[string]interface{})
		data["User"] = u
		// Render form to fix
		templates.Render(w, r, "user/new", data)
	}
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := user.User{
		Email:    r.PostForm.Get("user[email]"),
		Password: r.PostForm.Get("user[password]"),
	}
	if u.Login() {
		helpers.SignIn(r, u)
		flash.Success(r, "Successfully logged in!")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// Erase password for good practice.
		u.Password = ""
		data := make(map[string]interface{})
		data["User"] = u
		// Render form to fix
		templates.Render(w, r, "user/sign_in", data)
	}
}

func LogoutController(w http.ResponseWriter, r *http.Request) {
	helpers.SignOut(r)
	flash.Success(r, "Successfully signed out.")
	http.Redirect(w, r, "/", http.StatusFound)
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
}
