// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: feeds.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING id, name, url, user_id, created_at, updated_at
`

type CreateFeedParams struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Url    string    `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT id, name, url, user_id, created_at, updated_at FROM feeds
`

func (q *Queries) GetFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
