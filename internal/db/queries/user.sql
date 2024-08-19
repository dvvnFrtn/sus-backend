-- name: AddUser :execresult
INSERT INTO users (id, email, password, oauth_id, name, role)
VALUES (?, ?, ?, ?, ?, ?);