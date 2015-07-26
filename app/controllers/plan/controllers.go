package plan

import (
	"github.com/ShaneBurkhart/PlanSource/models/plan"
	"github.com/ShaneBurkhart/PlanSource/util"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ShowController(w http.ResponseWriter, r *http.Request) {
	// TODO some kind of validation that the plan belongs to the user.
	// TODO if the plan doesn't exist, return 404
	id := getId(r)
	d := make(map[string]interface{})
	d["plan"] = plan.Find(id)
	w.Write(util.JSON(d))
}

func CreateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := plan.Plan{
		Name:  r.PostForm.Get("plan[name]"),
		Num:   getFormInt("plan[num]", r),
		JobId: getFormInt("plan[job_id]", r),
	}
	d := make(map[string]interface{})
	if p.Save() {
		d["plan"] = p
	}
	w.Write(util.JSON(d))
}

func UpdateController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := plan.Find(getId(r))
	updateAttributes(p, r)
	p.Save()
	d := make(map[string]interface{})
	if p.Save() {
		d["job"] = p
	}
	w.Write(util.JSON(d))
}

func DeleteController(w http.ResponseWriter, r *http.Request) {
	// TODO 404 if no job
	p := plan.Find(getId(r))
	p.Delete()
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

func getFormInt(fieldName string, r *http.Request) int {
	// TODO this error handling will probably need to be different.
	i, err := strconv.Atoi(r.PostForm.Get(fieldName))
	util.CheckErr(err)
	return i
}

func updateAttributes(p *plan.Plan, r *http.Request) {
	// TODO this needs to be changed.  Not a good place to put this.
	if name := r.PostForm.Get("plan[name]"); name != "" {
		p.Name = name
	}
	if num := getFormInt("plan[num]", r); num > 0 {
		p.Num = num
	}
	if jobId := getFormInt("plan[job_id]", r); jobId > 0 {
		p.JobId = jobId
	}
}
