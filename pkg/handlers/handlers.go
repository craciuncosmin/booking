package handlers

import (
	"net/http"

	"github.com/craciuncosmin/booking/pkg/config"
	"github.com/craciuncosmin/booking/pkg/models"
	"github.com/craciuncosmin/booking/pkg/render"
)

//-----------------------------------------------------

// gotta read about repos.. really don't get this code
// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//-----------------------------------------------------

// in order for a function to respond to a request from a web browser, it has to handle 2 parameters:
// a response writer (w) and a request (r*)

// Home is the home page handler and has access to the Repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler and has access to the Repository
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "suhpreyez mothafucka"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
