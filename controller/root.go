package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/service"
)

type rootPage struct {
	Title string
	Body  string
}

// RegisterRootController add the "/" route and initializes constants for routes controlled by the "Root" controller
func RegisterRootController(server service.Server, config service.Config) {
	t := readTemplates(config)

	server.Get("/", welcomePage(t))
}

func readTemplates(config service.Config) *template.Template {
	t, err := template.New("index.html").ParseFiles(config.GetString("views_index"))

	if err != nil {
		panic(fmt.Errorf("Could not read index template: %s \n", err))
	}

	return t
}

func welcomePage(t *template.Template) func(model.Context, http.ResponseWriter, *http.Request) {
	return func(c model.Context, w http.ResponseWriter, r *http.Request) {
		page := rootPage{Title: "This Is The Title", Body: "This is the Body"}
		err := t.ExecuteTemplate(w, "index.html", page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
