// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: notes.sql

package database

import (
	"context"
)

const createNote = `-- name: CreateNote :one
insert into notes(id, created_at, updated_at, note)
values(null, datetime(), datetime(), ?)
returning id, created_at, updated_at, note
`

func (q *Queries) CreateNote(ctx context.Context, note string) (Note, error) {
	row := q.db.QueryRowContext(ctx, createNote, note)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Note,
	)
	return i, err
}

const getAllNotes = `-- name: GetAllNotes :many
select id, created_at, updated_at, note from notes
order by datetime(created_at) desc
`

func (q *Queries) GetAllNotes(ctx context.Context) ([]Note, error) {
	rows, err := q.db.QueryContext(ctx, getAllNotes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Note
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Note,
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
