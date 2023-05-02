package handlers

import (
	"net/http"

	"github.com/hiimdjango/gostarter/pkg/render"
)

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
