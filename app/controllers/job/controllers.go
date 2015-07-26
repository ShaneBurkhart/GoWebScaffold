package job

import (
	"github.com/ShaneBurkhart/PlanSource/helpers"
	"github.com/ShaneBurkhart/PlanSource/models/job"
	"github.com/ShaneBurkhart/PlanSource/models/user"
	"github.com/ShaneBurkhart/PlanSource/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	d := make(map[string]interface{})
	jobs := job.FindByUserId(helpers.CurrentUser(r).Id)

	c := make(chan *user.User)
	for _, j := range jobs {
		go func(userId int) {
			c <- user.Find(userId)
		}(j.UserId)
	}

	// TODO Dirty as all hell.  Probably want to clean up.
	// Right now, it works.
	users := make([]*user.User, 0)
	for _ = range jobs {
		u := <-c
		if u == nil {
			continue
		}
		var b bool
		for _, v := range users {
			if u.Id == v.Id {
				b = true
				break
			}
		}
		if !b {
			users = append(users, u)
		}
	}

	d["jobs"] = jobs
	d["users"] = users
	w.Write(util.JSON(d))
}

func ShowController(w http.ResponseWriter, r *http.Request) {
	// TODO if the job doesn't exist, return 404
	id := getId(r)
	d := make(map[string]interface{})
	d["job"] = job.Find(id, helpers.CurrentUser(r).Id)
	w.Write(util.JSON(d))
}

func CreateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	j := job.Job{
		Name:   r.PostForm.Get("job[name]"),
		UserId: helpers.CurrentUser(r).Id,
	}
	d := make(map[string]interface{})
	if j.Save() {
		d["job"] = j
	}
	w.Write(util.JSON(d))
}

func UpdateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	j := job.Find(getId(r), helpers.CurrentUser(r).Id)
	updateAttributes(j, r)
	j.Save()
	d := make(map[string]interface{})
	if j.Save() {
		d["job"] = j
	}
	w.Write(util.JSON(d))
}

func DeleteController(w http.ResponseWriter, r *http.Request) {
	// TODO 404 if no job
	j := job.Find(getId(r), helpers.CurrentUser(r).Id)
	j.Delete()
	w.Write(util.JSON(make(map[string]interface{})))
}

func getId(r *http.Request) int {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	// This will return a 500 instead of 404 because we have parameter
	// validation in routes.go that only allows numbers for id.  If there
	// is an error here, it is something different and most likely a 500.
	util.CheckErr(err)
	return id
}

func updateAttributes(j *job.Job, r *http.Request) {
	// TODO this needs to be changed.  Not a good place to put this.
	if name := r.PostForm.Get("job[name]"); name != "" {
		j.Name = name
	}
}
