package http

import (
	"github.com/earlofurl/lpotl-go"
	"github.com/earlofurl/lpotl-go/sqlc"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (s *Server) getAllSeriesHandler(w http.ResponseWriter, r *http.Request) {
	series, err := s.seriesService.FindAllSeries(r.Context())
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, series)
}

func (s *Server) getSeriesHandler(w http.ResponseWriter, r *http.Request) {
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

	series, err := s.seriesService.FindSeries(r.Context(), int32(n))
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, series)
}

func (s *Server) createSeriesHandler(w http.ResponseWriter, r *http.Request) {
	var req lpotl.CreateSeriesRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.CreateSeriesParams{
		Name:      req.Name,
		PodcastID: req.PodcastID,
	}

	series, err := s.seriesService.CreateSeries(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusCreated, series)
}

func (s *Server) updateSeriesHandler(w http.ResponseWriter, r *http.Request) {
	var req lpotl.UpdateSeriesRequest
	err := req.Bind(r.Body)
	if err != nil {
		Json(w, http.StatusBadRequest, err)
		return
	}

	arg := &sqlc.UpdateSeriesParams{
		Name: req.Name,
		ID:   req.ID,
	}

	series, err := s.seriesService.UpdateSeries(r.Context(), arg)
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusOK, series)
}

func (s *Server) deleteSeriesHandler(w http.ResponseWriter, r *http.Request) {
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

	err = s.seriesService.DeleteSeries(r.Context(), int32(n))
	if err != nil {
		Json(w, http.StatusInternalServerError, err)
		return
	}
	Json(w, http.StatusNoContent, nil)
}
