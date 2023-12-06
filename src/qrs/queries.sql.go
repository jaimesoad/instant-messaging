// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: queries.sql

package qrs

import (
	"context"
	"database/sql"
)

const deletePending = `-- name: DeletePending :execresult
delete from
    Messages
where
    recepient_id = (
        select
            id
        from
            Users
        where
            username = ?
    )
`

func (q *Queries) DeletePending(ctx context.Context, username string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deletePending, username)
}

const existUser = `-- name: ExistUser :one
SELECT
    COUNT(1) > 0
from
    Users
where
    username = ?
`

func (q *Queries) ExistUser(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, existUser, username)
	var column_1 bool
	err := row.Scan(&column_1)
	return column_1, err
}

const findUsers = `-- name: FindUsers :many
select username
from Users
where username like ? || '%'
limit 10
`

func (q *Queries) FindUsers(ctx context.Context, username string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, findUsers, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		items = append(items, username)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT
    id, username
from
    Users
where
    username = ?
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const newUser = `-- name: NewUser :execresult
INSERT INTO
    Users (username)
VALUES
    (?)
`

func (q *Queries) NewUser(ctx context.Context, username string) (sql.Result, error) {
	return q.db.ExecContext(ctx, newUser, username)
}

const pendingMessages = `-- name: PendingMessages :one
select
    count(1) > 0
from
    Messages
where
    recepient_id = (
        select
            id
        from
            Users
        where
            username = ?
    )
`

func (q *Queries) PendingMessages(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, pendingMessages, username)
	var column_1 bool
	err := row.Scan(&column_1)
	return column_1, err
}

const saveMessage = `-- name: SaveMessage :execresult
insert into
    Messages (content, recepient_id, sender_id)
VALUES
    (
        ?,
        (
            select
                id
            from
                Users as t1
            where
                t1.username = ?2 -- name: Recepient
        ),
        (
            select
                id
            from
                Users as t2
            where
                t2.username = ?3 -- name: Sender
        )
    )
`

type SaveMessageParams struct {
	Content   string `json:"content"`
	Recepient string `json:"recepient"`
	Sender    string `json:"sender"`
}

func (q *Queries) SaveMessage(ctx context.Context, arg SaveMessageParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, saveMessage, arg.Content, arg.Recepient, arg.Sender)
}

const savedMessages = `-- name: SavedMessages :many
SELECT
    msg.content,
    (select t1.username from Users as t1 where t1.id = msg.sender_id) as sender
from
    Users as u,
    Messages as msg
where
    u.username = ?
    AND u.id = msg.recepient_id
`

type SavedMessagesRow struct {
	Content string `json:"content"`
	Sender  string `json:"sender"`
}

func (q *Queries) SavedMessages(ctx context.Context, username string) ([]SavedMessagesRow, error) {
	rows, err := q.db.QueryContext(ctx, savedMessages, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SavedMessagesRow
	for rows.Next() {
		var i SavedMessagesRow
		if err := rows.Scan(&i.Content, &i.Sender); err != nil {
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
