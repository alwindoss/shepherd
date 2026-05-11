package server

import (
	"github.com/go-chi/chi/v5"
)

func configureMemberRoutes(mux *chi.Mux, hdlr ResourceHandler) {
	mux.Route("/members", func(r chi.Router) {
		r.Get("/", hdlr.Index)
		r.Get("/new", hdlr.New)
		r.Post("/", hdlr.Create)
		r.Get("/{id}", hdlr.Show)
		r.Get("/{id}/edit", hdlr.Edit)
		r.Patch("/{id}", hdlr.UpdateFull)
		r.Put("/{id}", hdlr.UpdatePartial)
		r.Delete("/{id}", hdlr.Delete)
	})
}

func configureEventRoutes(mux *chi.Mux, hdlr ResourceHandler) {
	mux.Route("/events", func(r chi.Router) {
		r.Get("/", hdlr.Index)
		r.Get("/new", hdlr.New)
		r.Post("/", hdlr.Create)
		r.Get("/{id}", hdlr.Show)
		r.Get("/{id}/edit", hdlr.Edit)
		r.Patch("/{id}", hdlr.UpdateFull)
		r.Put("/{id}", hdlr.UpdatePartial)
		r.Delete("/{id}", hdlr.Delete)
	})
}

func configureDonationsRoutes(mux *chi.Mux, hdlr ResourceHandler) {
	mux.Route("/donations", func(r chi.Router) {
		r.Get("/", hdlr.Index)
		r.Get("/new", hdlr.New)
		r.Post("/", hdlr.Create)
		r.Get("/{id}", hdlr.Show)
		r.Get("/{id}/edit", hdlr.Edit)
		r.Patch("/{id}", hdlr.UpdateFull)
		r.Put("/{id}", hdlr.UpdatePartial)
		r.Delete("/{id}", hdlr.Delete)
	})
}
