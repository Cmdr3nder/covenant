package controller

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/service"
)

// RegisterStaticFileController adds routes and initializes constants for routes controlled by the "Resume" controller
func RegisterStaticFileController(server service.Server, config service.Config) {
	cssRegex := service.GetRouteBuilder().AppendPart("css").AppendPart("(?P<fileName>.+\\.css)")
	server.Get(cssRegex.MustCompile(), getCSSFile(config))

	jsRegex := service.GetRouteBuilder().AppendPart("js").AppendPart("(?P<fileName>.+\\.js)")
	server.Get(jsRegex.MustCompile(), getJavaScriptFile(config))

	server.Get("/favicon.ico", getFavicon(config))
}

func getCSSFile(config service.Config) func(model.Context, http.ResponseWriter, *http.Request) {
	cssDir := config.GetString("css")

	return func(c model.Context, w http.ResponseWriter, r *http.Request) {
		sendFile(w, "text/css", cssDir+"/"+c.GetURLParam("fileName"), config.GetBool("debug"))
	}
}

func getJavaScriptFile(config service.Config) func(model.Context, http.ResponseWriter, *http.Request) {
	jsDir := config.GetString("js")

	return func(c model.Context, w http.ResponseWriter, r *http.Request) {
		sendFile(w, "application/javascript", jsDir+"/"+c.GetURLParam("fileName"), config.GetBool("debug"))
	}
}

func getFavicon(config service.Config) func(model.Context, http.ResponseWriter, *http.Request) {
	faviconPath := config.GetString("favicon")

	return func(c model.Context, w http.ResponseWriter, r *http.Request) {
		sendFile(w, "image/x-icon", faviconPath, config.GetBool("debug"))
	}
}

func sendFile(w http.ResponseWriter, contentType string, path string, debug bool) {
	w.Header().Set("Content-type", contentType)

	fileBytes, err := getFile(path, debug)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	buffer := bytes.NewBuffer(fileBytes)

	if _, err := buffer.WriteTo(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var files = make(map[string][]byte)

func getFile(path string, debug bool) ([]byte, error) {
	if files[path] == nil || debug {
		fileBytes, err := ioutil.ReadFile(path)

		if err != nil {
			return nil, err
		}

		files[path] = fileBytes
	}

	return files[path], nil
}
