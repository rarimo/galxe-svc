package service

import (
	"github.com/go-chi/chi"
	"github.com/rarimo/galxe-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/integrations/galxe-svc", func(r chi.Router) {
		r.Get("/info", handlers.InfoAddress)
	})

	return r
}
