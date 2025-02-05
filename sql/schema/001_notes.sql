-- +goose Up

create table notes(
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL,
	note TEXT NOT NULL
);


-- +goose Down
drop table notes;
