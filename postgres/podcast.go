package postgres

import (
	"context"
	"github.com/earlofurl/lpotl-go"
	"github.com/earlofurl/lpotl-go/sqlc"
)

// ensure service implements interface
var _ lpotl.PodcastService = (*PodcastService)(nil)

type PodcastService struct {
	store sqlc.Store
}

func NewPodcastService(store *sqlc.Store) *PodcastService {
	return &PodcastService{store: *store}
}

// FindAllPodcasts retrieves all Podcasts.
func (s *PodcastService) FindAllPodcasts(ctx context.Context) ([]*sqlc.Podcast, error) {
	u, err := s.store.ListPodcasts(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindPodcast retrieves an Podcast by ID.
// Returns ENOTFOUND if Podcast does not exist.
func (s *PodcastService) FindPodcast(ctx context.Context, id int32) (*sqlc.Podcast, error) {
	u, err := s.store.GetPodcast(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, err
	}
	return u, nil
}

// CreatePodcast creates a new Podcast.
func (s *PodcastService) CreatePodcast(ctx context.Context, name string) (*sqlc.Podcast, error) {
	u, err := s.store.CreatePodcast(ctx, name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdatePodcast updates a Podcast.
func (s *PodcastService) UpdatePodcast(ctx context.Context, p *sqlc.UpdatePodcastParams) (*sqlc.Podcast, error) {
	u, err := s.store.UpdatePodcast(ctx, p)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// DeletePodcast deletes an Podcast by ID.
func (s *PodcastService) DeletePodcast(ctx context.Context, id int32) error {
	err := s.store.DeletePodcast(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
