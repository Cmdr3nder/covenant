package controller

import (
	"fmt"
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/service"
)

// RegisterStaticFileController adds routes and initializes constants for routes controlled by the "Resume" controller
func RegisterStaticFileController(server service.Server) {
	cssRegex := service.GetRouteBuilder().AppendPart("css").AppendPart("(?P<fileName>.+\\.css)")
	server.Get(cssRegex.MustCompile(), getCSSFile)

	jsRegex := service.GetRouteBuilder().AppendPart("js").AppendPart("(?P<fileName>.+\\.js)")
	server.Get(jsRegex.MustCompile(), getJavaScriptFile)

	server.Get("/favicon.ico", getFavicon)
}

func getCSSFile(c model.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rendering \"css\" %s %s", c.GetURLParam("fileName"), r.FormValue("kingdom"))
}

func getJavaScriptFile(c model.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rendering \"js\" %s", c.GetURLParam("fileName"))
}

func getFavicon(c model.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Oops Favicon Not Found Lolz")
}
