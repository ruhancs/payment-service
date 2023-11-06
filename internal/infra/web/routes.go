package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/health"))

	mux.Route("/payment",func(r chi.Router) {
		mux.Use(app.Auth)
		mux.Post("/payment", app.PaymentHandler)
	})

	return mux
}
