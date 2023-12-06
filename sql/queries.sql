-- name: SavedMessages :many
SELECT
    msg.content,
    (select t1.username from Users as t1 where t1.id = msg.sender_id) as sender
from
    Users as u,
    Messages as msg
where
    u.username = ?
    AND u.id = msg.recepient_id;

-- name: NewUser :execresult
INSERT INTO
    Users (username)
VALUES
    (?);

-- name: ExistUser :one
SELECT
    COUNT(1) > 0
from
    Users
where
    username = ?;

-- name: GetUser :one
SELECT
    *
from
    Users
where
    username = ?;

-- name: SaveMessage :execresult
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
                t1.username = sqlc.arg(recepient) -- name: Recepient
        ),
        (
            select
                id
            from
                Users as t2
            where
                t2.username = sqlc.arg(sender) -- name: Sender
        )
    );

-- name: PendingMessages :one
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
    );

-- name: DeletePending :execresult
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
    );

-- name: FindUsers :many
select username
from Users
where username like ? || '%'
limit 10;