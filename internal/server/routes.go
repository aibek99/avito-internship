package server

import (
	"avito-internship/internal/api/employee"

	"github.com/go-chi/chi/v5"
)

func MapRoutes(group chi.Router, h employee.Handlers) {
	group.Route("/api", func(r chi.Router) {
		r.Get("/ping", h.Ping())
		//r.Get("/tenders", h.Tenders())
		//r.Delete("/turn_in/{orderID}", h.TurnInOrder())
	})
}
