package postgres

import (
	"context"
	"github.com/earlofurl/lpotl-go"
	"github.com/earlofurl/lpotl-go/sqlc"
)

// ensure service implements interface
var _ lpotl.EpisodeService = (*EpisodeService)(nil)

type EpisodeService struct {
	store sqlc.Store
}

func NewEpisodeService(store *sqlc.Store) *EpisodeService {
	return &EpisodeService{store: *store}
}

// FindAllEpisodes retrieves all Episodes.
func (s *EpisodeService) FindAllEpisodes(ctx context.Context) ([]*sqlc.Episode, error) {
	u, err := s.store.ListEpisodes(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindEpisode retrieves an Episode by ID.
// Returns ENOTFOUND if Episode does not exist.
func (s *EpisodeService) FindEpisode(ctx context.Context, id int32) (*sqlc.Episode, error) {
	u, err := s.store.GetEpisode(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, err
	}
	return u, nil
}

// CreateEpisode creates a new Episode.
func (s *EpisodeService) CreateEpisode(ctx context.Context, e *sqlc.CreateEpisodeParams) (*sqlc.Episode, error) {
	u, err := s.store.CreateEpisode(ctx, e)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateEpisode updates an Episode.
func (s *EpisodeService) UpdateEpisode(ctx context.Context, e *sqlc.UpdateEpisodeParams) (*sqlc.Episode, error) {
	u, err := s.store.UpdateEpisode(ctx, e)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// DeleteEpisode deletes an Episode by ID.
func (s *EpisodeService) DeleteEpisode(ctx context.Context, id int32) error {
	err := s.store.DeleteEpisode(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
