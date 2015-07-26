package templates

import (
	"fmt"
	"github.com/ShaneBurkhart/PlanSource/flash"
	"github.com/ShaneBurkhart/PlanSource/helpers"
	"github.com/dustin/go-humanize"
	"html/template"
	"net/http"
	"path"
	"strings"
)

func Render(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}

	data["Controller"] = getController(name)
	data["Flashes"] = flash.Flashes(r)
	data["CurrentUser"] = helpers.CurrentUser(r)

	p := path.Join("templates", fmt.Sprintf("%s.html", name))
	t, err := template.New("t").Funcs(template.FuncMap{
		"Title":     strings.Title,
		"HumanTime": humanize.Time,
	}).ParseFiles(
		p,
		"templates/layouts/application.html",
		"templates/layouts/navigation.html",
		"templates/layouts/messages.html",
		"templates/helpers/form.html",
	)
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "base", data)
	if err != nil {
		panic(err)
	}
}

func getController(name string) string {
	return name[:strings.Index(name, "/")]
}
