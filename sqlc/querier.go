// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateEpisode(ctx context.Context, arg *CreateEpisodeParams) (*Episode, error)
	CreatePodcast(ctx context.Context, name string) (*Podcast, error)
	CreateSeries(ctx context.Context, arg *CreateSeriesParams) (*Series, error)
	DeleteEpisode(ctx context.Context, id int32) error
	DeletePodcast(ctx context.Context, id int32) error
	DeleteSeries(ctx context.Context, id int32) error
	GetEpisode(ctx context.Context, id int32) (*Episode, error)
	GetPodcast(ctx context.Context, id int32) (*Podcast, error)
	GetSeries(ctx context.Context, id int32) (*Series, error)
	ListEpisodes(ctx context.Context) ([]*Episode, error)
	ListPodcasts(ctx context.Context) ([]*Podcast, error)
	ListSeries(ctx context.Context) ([]*Series, error)
	SearchEpisodes(ctx context.Context, websearchToTsquery string) ([]*SearchEpisodesRow, error)
	UpdateEpisode(ctx context.Context, arg *UpdateEpisodeParams) (*Episode, error)
	UpdatePodcast(ctx context.Context, arg *UpdatePodcastParams) (*Podcast, error)
	UpdateSeries(ctx context.Context, arg *UpdateSeriesParams) (*Series, error)
}

var _ Querier = (*Queries)(nil)
