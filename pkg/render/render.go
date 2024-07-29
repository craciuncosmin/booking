package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/craciuncosmin/booking/pkg/config"
	"github.com/craciuncosmin/booking/pkg/models"
)

var app *config.AppConfig

// NewTemplates allows for the use of the global config in the config package inside the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// create a template cache by getting it from the app config
		tc = app.TemplateCache
	} else {
		// rebuild the template cache
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("template not found in cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	//when rendering a HTML template that uses a layout, you must have the template you want to render as the first thing
	//that means *.page.tmpl first
	//then base.layout.tmpl

	// get all files named *.page.tmpl from templates dir
	htmlPages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	// range through all HTML pages (all files named "./templates/name.page.tmpl")
	for _, page := range htmlPages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
		//log.Println(cache)
	}

	return cache, nil
}
