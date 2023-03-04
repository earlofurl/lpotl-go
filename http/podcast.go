package http

import (
	"github.com/earlofurl/lpotl-go"
	"github.com/earlofurl/lpotl-go/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (s *Server) getAllPodcastsHandler(w http.ResponseWriter, r *http.Request) {
	Podcasts, err := s.podcastService.FindAllPodcasts(r.Context())
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, Podcasts)
}

func (s *Server) getPodcastHandler(w http.ResponseWriter, r *http.Request) {
	i := chi.URLParam(r, "id")
	if i == "" {
		Json(w, http.StatusBadRequest, nil)
		return
	}

	n, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	if n < 1 {
		Json(w, http.StatusBadRequest, err)
		return
	}

	podcast, err := s.podcastService.FindPodcast(r.Context(), int32(n))
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, podcast)
}

func (s *Server) createPodcastHandler(w http.ResponseWriter, r *http.Request) {
	var req lpotl.CreatePodcastRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Msgf("createPodcastHandler: req.Name: %s", req.Name)

	podcast, err := s.podcastService.CreatePodcast(r.Context(), req.Name)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusCreated, podcast)
}

func (s *Server) updatePodcastHandler(w http.ResponseWriter, r *http.Request) {
	var req lpotl.UpdatePodcastRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdatePodcastParams{
		Name: req.Name,
		ID:   req.ID,
	}

	podcast, err := s.podcastService.UpdatePodcast(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, podcast)
}

func (s *Server) deletePodcastHandler(w http.ResponseWriter, r *http.Request) {
	i := chi.URLParam(r, "id")
	if i == "" {
		Json(w, http.StatusBadRequest, nil)
		return
	}

	n, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}
	if n < 1 {
		Json(w, http.StatusBadRequest, err)
		return
	}

	err = s.podcastService.DeletePodcast(r.Context(), int32(n))
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusNoContent, nil)
}
