package main

import (
	"net/http"

	"github.com/craciuncosmin/booking/pkg/handlers"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	//middleware is code that can handle dynamic requests/events
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
