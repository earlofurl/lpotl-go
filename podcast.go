package lpotl

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/earlofurl/lpotl-go/sqlc"
	"io"
)

type PodcastService interface {
	FindAllPodcasts(ctx context.Context) ([]*sqlc.Podcast, error)
	FindPodcast(ctx context.Context, id int32) (*sqlc.Podcast, error)
	CreatePodcast(ctx context.Context, f string) (*sqlc.Podcast, error)
	UpdatePodcast(ctx context.Context, f *sqlc.UpdatePodcastParams) (*sqlc.Podcast, error)
	DeletePodcast(ctx context.Context, id int32) error
}

type CreatePodcastRequest struct {
	Name string `json:"name"`
}

type UpdatePodcastRequest struct {
	Name sql.NullString `json:"name"`
	ID   int32          `json:"id"`
}

func (r *CreatePodcastRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdatePodcastRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
