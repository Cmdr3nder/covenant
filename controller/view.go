package controller

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "path/filepath"
    "net/http"

    "github.com/zenazn/goji"
)

func RegisterViewController() {
    goji.Get("/view/:view", viewHandler)
}

func viewHandler(response http.ResponseWriter, request *http.Request) {
    title := request.URL.Path[len("/view/"):]
    page, err := loadPage(title)
    if err == nil {
        fmt.Fprintf(response, page.Html())
    } else {
        webError(response, err)
    }
}

func webError(response http.ResponseWriter, err error) {
    // errPage := Page{Title: "Internal Server Error", Body: []byte(err.Error())}
    // fmt.Fprintf(response, errPage.Html())
    http.Error(response, err.Error(), http.StatusInternalServerError)
}

//Page "Class"
type Page struct {
    Title string
    Body []byte
}

func (p *Page) save() error {
    filename, _ := fileName(p.Title)
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func (p *Page) Html() string {
    html := bytes.Buffer{}
    html.WriteString("<html>")
    html.WriteString("<head>")
    html.WriteString("<title>")
    html.WriteString(p.Title)
    html.WriteString("</title>")
    html.WriteString("</head>")
    html.WriteString("<body>")
    html.WriteString("<h1>")
    html.WriteString(p.Title)
    html.WriteString("</h1>")
    html.WriteString("<div>")
    html.Write(p.Body)
    html.WriteString("</div>")
    html.WriteString("</body>")
    html.WriteString("</html>")
    return html.String()
}

func loadPage(title string) (*Page, error) {
    filename, err := fileName(title)
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

//Utility Functions
func fileName(title string) (string, error) {
    file_path := bytes.Buffer{}
    file_path.WriteRune('.')
    file_path.WriteRune(filepath.Separator)
    file_path.WriteString("pages")
    makeFolder(file_path.String()) //Make Sure Folder Exists
    file_path.WriteRune(filepath.Separator)
    file_path.WriteString(title)
    file_path.WriteString(".txt")
    return filepath.Abs(file_path.String())
}

func makeFolder(path string) error {
    return nil
}
