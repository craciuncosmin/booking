package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/craciuncosmin/booking/pkg/handlers"
	"github.com/craciuncosmin/booking/pkg/render"

	"github.com/craciuncosmin/booking/pkg/config"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8877"

var app config.AppConfig

var session *scs.SessionManager

func main() {
	//change this to true when in prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 20 * time.Minute
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	//-----------------------------------------------------

	//gotta read about repos.. really don't get this code
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//-----------------------------------------------------

	render.NewTemplates(&app)

	fmt.Println("Starting application on port", portNumber)

	serving := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = serving.ListenAndServe()
	log.Fatal(err)
}
