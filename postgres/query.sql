-- name: GetEpisode :one
SELECT *
FROM episodes
WHERE id = $1
LIMIT 1;

-- name: ListEpisodes :many
SELECT *
FROM episodes
ORDER BY name;

-- name: CreateEpisode :one
INSERT INTO episodes (name, number_series, number_overall, release_date, description, body, transcript_url, podcast_id,
                      series_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateEpisode :one
UPDATE episodes
SET name           = COALESCE(sqlc.narg(name), name),
    number_series  = COALESCE(sqlc.narg(number_series), number_series),
    number_overall = COALESCE(sqlc.narg(number_overall), number_overall),
    release_date   = COALESCE(sqlc.narg(release_date), release_date),
    description    = COALESCE(sqlc.narg(description), description),
    body           = COALESCE(sqlc.narg(body), body),
    transcript_url = COALESCE(sqlc.narg(transcript_url), transcript_url),
    podcast_id     = COALESCE(sqlc.narg(podcast_id), podcast_id),
    series_id      = COALESCE(sqlc.narg(series_id), series_id),
    last_updated   = NOW()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: SearchEpisodes :many
SELECT *,
       ts_headline(body, websearch_to_tsquery($1))::text AS headline
FROM episodes
WHERE fts_doc_en @@ websearch_to_tsquery($1);

-- name: DeleteEpisode :exec
DELETE
FROM episodes
WHERE id = $1;

-- name: GetPodcast :one
SELECT *
FROM podcasts
WHERE id = $1
LIMIT 1;

-- name: ListPodcasts :many
SELECT *
FROM podcasts
ORDER BY name;

-- name: CreatePodcast :one
INSERT INTO podcasts (name)
VALUES ($1)
RETURNING *;

-- name: UpdatePodcast :one
UPDATE podcasts
SET name         = COALESCE(sqlc.narg(name), name),
    last_updated = NOW()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeletePodcast :exec
DELETE
FROM podcasts
WHERE id = $1;

-- name: GetSeries :one
SELECT *
FROM series
WHERE id = $1
LIMIT 1;

-- name: ListSeries :many
SELECT *
FROM series
ORDER BY name;

-- name: CreateSeries :one
INSERT INTO series (name, podcast_id)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateSeries :one
UPDATE series
SET name         = COALESCE(sqlc.narg(name), name),
    podcast_id   = COALESCE(sqlc.narg(podcast_id), podcast_id),
    last_updated = NOW()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteSeries :exec
DELETE
FROM series
WHERE id = $1;