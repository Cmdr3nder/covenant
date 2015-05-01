package controller

import (
    "net/http"
    "html/template"

    "github.com/ender4021/covenant/model"
    "github.com/ender4021/covenant/service"
)

var t = template.Must(template.New("name").Parse("view/shared/index.html"))

func RegisterRootController(server service.Server) {
    server.Get("/", welcomePage)
}

func welcomePage(c model.Context, w http.ResponseWriter, r *http.Request) {
    page := Page{Title: "This Is The Title", Body: []byte("This is the Body")}
    err := t.ExecuteTemplate(w, "name", page)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
