package http

import (
	"github.com/earlofurl/lpotl-go"
	"github.com/earlofurl/lpotl-go/sqlc"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (s *Server) getAllEpisodesHandler(w http.ResponseWriter, r *http.Request) {
	episodes, err := s.episodeService.FindAllEpisodes(r.Context())
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, episodes)
}

func (s *Server) getEpisodeHandler(w http.ResponseWriter, r *http.Request) {
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

	episode, err := s.episodeService.FindEpisode(r.Context(), int32(n))
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, episode)
}

func (s *Server) createEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	var req lpotl.CreateEpisodeRequest
	err := req.Bind(r.Body)
	if err != nil {
		log.Error().Msgf("createEpisodeHandler: req.Bind(r.Body) failed: %s", err)
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.CreateEpisodeParams{
		Name:          req.Name,
		NumberSeries:  req.NumberSeries,
		NumberOverall: req.NumberOverall,
		ReleaseDate:   req.ReleaseDate,
		Description:   req.Description,
		Body:          req.Body,
		TranscriptUrl: req.TranscriptUrl,
		PodcastID:     req.PodcastID,
		SeriesID:      req.SeriesID,
	}

	episode, err := s.episodeService.CreateEpisode(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusCreated, episode)
}

func (s *Server) updateEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	var req lpotl.UpdateEpisodeRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdateEpisodeParams{
		Name:          req.Name,
		NumberSeries:  req.NumberSeries,
		NumberOverall: req.NumberOverall,
		ReleaseDate:   req.ReleaseDate,
		Description:   req.Description,
		Body:          req.Body,
		TranscriptUrl: req.TranscriptUrl,
		PodcastID:     req.PodcastID,
		SeriesID:      req.SeriesID,
		ID:            req.ID,
	}

	item, err := s.episodeService.UpdateEpisode(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, item)
}

func (s *Server) deleteEpisodeHandler(w http.ResponseWriter, r *http.Request) {
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

	err = s.episodeService.DeleteEpisode(r.Context(), int32(n))
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusNoContent, nil)
}
