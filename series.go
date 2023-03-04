package lpotl

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/earlofurl/lpotl-go/sqlc"
	"io"
)

type SeriesService interface {
	FindAllSeries(ctx context.Context) ([]*sqlc.Series, error)
	FindSeries(ctx context.Context, id int32) (*sqlc.Series, error)
	CreateSeries(ctx context.Context, f *sqlc.CreateSeriesParams) (*sqlc.Series, error)
	UpdateSeries(ctx context.Context, f *sqlc.UpdateSeriesParams) (*sqlc.Series, error)
	DeleteSeries(ctx context.Context, id int32) error
}

type CreateSeriesRequest struct {
	Name      string `json:"name"`
	PodcastID int32  `json:"podcast_id"`
}

type UpdateSeriesRequest struct {
	Name      sql.NullString `json:"name"`
	PodcastID sql.NullInt32  `json:"podcast_id"`
	ID        int32          `json:"id"`
}

func (r *CreateSeriesRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdateSeriesRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
