package config

import (
	"fmt"
	"github.com/ShaneBurkhart/PlanSource/app"
	"github.com/ShaneBurkhart/PlanSource/controllers/admin"
	"github.com/ShaneBurkhart/PlanSource/controllers/job"
	"github.com/ShaneBurkhart/PlanSource/controllers/marketing"
	"github.com/ShaneBurkhart/PlanSource/controllers/mobile"
	"github.com/ShaneBurkhart/PlanSource/controllers/plan"
	"github.com/ShaneBurkhart/PlanSource/controllers/user"
	"github.com/ShaneBurkhart/PlanSource/flash"
	"github.com/ShaneBurkhart/PlanSource/helpers"
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func SetupRoutes(r *mux.Router) {
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Handle("/users", authAdminUser(admin.UserIndexController)).Methods("GET")
	adminRouter.Handle("/users/{id:[0-9]+}/edit", authAdminUser(admin.UserEditController)).Methods("GET")
	adminRouter.Handle("/users/{id:[0-9]+}/edit", authAdminUser(admin.UserUpdateController)).Methods("POST")
	adminRouter.Handle("/users/{id:[0-9]+}/delete", authAdminUser(admin.UserDeleteController)).Methods("POST")

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Handle("/jobs", authUser(job.IndexController)).Methods("GET")
	apiRouter.Handle("/jobs/{id:[0-9]+}", authUser(job.ShowController)).Methods("GET")
	apiRouter.Handle("/jobs", authUser(job.CreateController)).Methods("POST")
	apiRouter.Handle("/jobs/{id:[0-9]+}", authUser(job.UpdateController)).Methods("PUT")
	apiRouter.Handle("/jobs/{id:[0-9]+}", authUser(job.DeleteController)).Methods("DELETE")
	apiRouter.Handle("/plans/{id:[0-9]+}", authUser(plan.ShowController)).Methods("GET")
	apiRouter.Handle("/plans", authUser(plan.CreateController)).Methods("POST")
	apiRouter.Handle("/plans/{id:[0-9]+}", authUser(plan.UpdateController)).Methods("PUT")
	apiRouter.Handle("/plans/{id:[0-9]+}", authUser(plan.DeleteController)).Methods("DELETE")

	r.Handle("/app", authUser(app.AppController)).Methods("GET")

	r.Handle("/user", redirectUser(user.CreateController)).Methods("POST")
	r.Handle("/user/edit", authUser(user.EditController)).Methods("GET")
	r.Handle("/user/edit", authUser(user.UpdateController)).Methods("POST")
	r.Handle("/user/sign_in", redirectUser(user.LoginController)).Methods("POST")
	// Ideally we want this to be a DELETE method but it is a bitch to send a DELETE
	// request client side.  POST will do for now.
	r.Handle("/user/sign_out", authUser(user.LogoutController)).Methods("POST")

	r.HandleFunc("/mobile", mobile.MobileController).Methods("GET")
	r.Handle("/", redirectUser(marketing.HomeController)).Methods("GET")
}

func Serve(r *mux.Router) {
	store := cookiestore.New([]byte("some-super-secret-key"))
	store.Options(sessions.Options{
		Path:     "/",      // to match all requests
		MaxAge:   3600 * 8, // 8 hour
		HTTPOnly: true,
	})

	n := negroni.Classic()
	n.Use(sessions.Sessions("plansource", store))
	n.Use(negroni.HandlerFunc(UpdateLastSeen))

	n.UseHandler(r)
	n.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

// Use this to redirect user from pages like the homepage to app page when they are
// already signed in.
type redirectUser func(http.ResponseWriter, *http.Request)

func (fn redirectUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if helpers.CurrentUser(r) != nil {
		// Send them to the app page since they are signed in.
		http.Redirect(w, r, "/app", http.StatusFound)
	} else {
		fn(w, r)
	}
}

type authUser func(http.ResponseWriter, *http.Request)

func (fn authUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if helpers.CurrentUser(r) != nil {
		fn(w, r)
	} else {
		// TODO Remember where the user was going so we can send them back to that page.
		// TODO Think about where to redirect.  Currently, we don't have sign in page.
		// We just have the home page.
		flash.Info(r, "You need to sign in before you can see that page.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

type authAdminUser func(http.ResponseWriter, *http.Request)

func (fn authAdminUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u := helpers.CurrentUser(r)
	if u != nil && u.IsAdmin() {
		fn(w, r)
	} else {
		// TODO 404 Redirect.
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
