package lpotl

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"time"

	"github.com/earlofurl/lpotl-go/sqlc"
)

type EpisodeService interface {
	FindAllEpisodes(ctx context.Context) ([]*sqlc.Episode, error)
	FindEpisode(ctx context.Context, id int32) (*sqlc.Episode, error)
	CreateEpisode(ctx context.Context, f *sqlc.CreateEpisodeParams) (*sqlc.Episode, error)
	UpdateEpisode(ctx context.Context, f *sqlc.UpdateEpisodeParams) (*sqlc.Episode, error)
	DeleteEpisode(ctx context.Context, id int32) error
}

type CreateEpisodeRequest struct {
	Name          string    `json:"name"`
	NumberSeries  int32     `json:"number_series"`
	NumberOverall int32     `json:"number_overall"`
	ReleaseDate   time.Time `json:"release_date"`
	Description   string    `json:"description"`
	Body          string    `json:"body"`
	TranscriptUrl string    `json:"transcript_url"`
	PodcastID     int32     `json:"podcast_id"`
	SeriesID      int32     `json:"series_id"`
}

type UpdateEpisodeRequest struct {
	Name          sql.NullString `json:"name"`
	NumberSeries  sql.NullInt32  `json:"number_series"`
	NumberOverall sql.NullInt32  `json:"number_overall"`
	ReleaseDate   sql.NullTime   `json:"release_date"`
	Description   sql.NullString `json:"description"`
	Body          sql.NullString `json:"body"`
	TranscriptUrl sql.NullString `json:"transcript_url"`
	PodcastID     sql.NullInt32  `json:"podcast_id"`
	SeriesID      sql.NullInt32  `json:"series_id"`
	ID            int32          `json:"id"`
}

func (r *CreateEpisodeRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdateEpisodeRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
