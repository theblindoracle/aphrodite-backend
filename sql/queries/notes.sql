
-- name: GetAllNotes :many
select * from notes
order by datetime(created_at) desc;

-- name: CreateNote :one
insert into notes(id, created_at, updated_at, note)
values(null, datetime(), datetime(), ?)
returning *;
