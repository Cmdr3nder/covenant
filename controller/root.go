package controller

import (
    "net/http"

    "html/template"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

var t = template.Must(template.New("name").Parse("view/shared/index.html"))

func RegisterRootController() {
    goji.Get("/", welcomePage)
}

func welcomePage(c web.C, w http.ResponseWriter, r *http.Request) {
    page := Page{Title: "This Is The Title", Body: []byte("This is the Body")}
    err := t.ExecuteTemplate(w, "name", page)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
