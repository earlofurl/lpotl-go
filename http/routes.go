package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s *Server) InitRoutes() {
	s.initVersion()
	//s.initHealth()
	//s.initSwagger()
	s.initEpisode()
	s.initPodcast()
	s.initSeries()
}

func (s *Server) initVersion() {
	s.router.Route("/version", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			Json(w, http.StatusOK, map[string]string{"version": s.Version})
		})
	})
}

func (s *Server) initEpisode() {
	s.router.Route("/api/episode", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getAllEpisodesHandler)
		router.Get("/{id}", s.getEpisodeHandler)
		router.Post("/", s.createEpisodeHandler)
		router.Put("/{id}", s.updateEpisodeHandler)
		router.Delete("/{id}", s.deleteEpisodeHandler)
	})
}

func (s *Server) initPodcast() {
	s.router.Route("/api/podcast", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getAllPodcastsHandler)
		router.Get("/{id}", s.getPodcastHandler)
		router.Post("/", s.createPodcastHandler)
		router.Put("/{id}", s.updatePodcastHandler)
		router.Delete("/{id}", s.deletePodcastHandler)
	})
}

func (s *Server) initSeries() {
	s.router.Route("/api/series", func(router chi.Router) {
		router.Use(JsonMiddleware)

		router.Get("/", s.getAllSeriesHandler)
		router.Get("/{id}", s.getSeriesHandler)
		router.Post("/", s.createSeriesHandler)
		router.Put("/{id}", s.updateSeriesHandler)
		router.Delete("/{id}", s.deleteSeriesHandler)
	})
}
