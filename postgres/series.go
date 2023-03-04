package postgres

import (
	"context"
	"github.com/earlofurl/lpotl-go"
	"github.com/earlofurl/lpotl-go/sqlc"
)

// ensure service implements interface
var _ lpotl.SeriesService = (*SeriesService)(nil)

type SeriesService struct {
	store sqlc.Store
}

func NewSeriesService(store *sqlc.Store) *SeriesService {
	return &SeriesService{store: *store}
}

// FindAllSeries retrieves all Series.
func (s *SeriesService) FindAllSeries(ctx context.Context) ([]*sqlc.Series, error) {
	u, err := s.store.ListSeries(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindSeries retrieves a Series by ID.
// Returns ENOTFOUND if Series does not exist.
func (s *SeriesService) FindSeries(ctx context.Context, id int32) (*sqlc.Series, error) {
	u, err := s.store.GetSeries(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, err
	}
	return u, nil
}

// CreateSeries creates a new Series.
func (s *SeriesService) CreateSeries(ctx context.Context, arg *sqlc.CreateSeriesParams) (*sqlc.Series, error) {
	u, err := s.store.CreateSeries(ctx, arg)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateSeries updates a Series.
func (s *SeriesService) UpdateSeries(ctx context.Context, arg *sqlc.UpdateSeriesParams) (*sqlc.Series, error) {
	u, err := s.store.UpdateSeries(ctx, arg)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// DeleteSeries deletes an Series by ID.
func (s *SeriesService) DeleteSeries(ctx context.Context, id int32) error {
	err := s.store.DeleteSeries(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
