package cmd

import(
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func CreateRouter() *chi.mux{
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return r
}