-- name: AddUser :execresult
INSERT INTO users (
    id, email, password, oauth_id, phone,
    name, role, img, is_premium, lvl,
    dob, institution, created_at, updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: EmailExists :one
SELECT COUNT(1) FROM users WHERE email = ?;

-- name: FindByEmail :one
SELECT * FROM users WHERE email = ?;
