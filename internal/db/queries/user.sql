-- name: AddUser :execresult
INSERT INTO users (
    id, email, password, oauth_id, phone,
    name, role, address, img, is_premium, lvl,
    dob, institution, created_at, updated_at
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: EmailExists :one
SELECT COUNT(1) FROM users WHERE email = ?;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = ?;

-- name: FindUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: UpdateUserByID :execresult
UPDATE users
SET username = ?, name = ?, address = ?, dob = ?, institution = ?
WHERE id = ?;

-- name: GetOrganizer :one
SELECT * FROM organizers WHERE id = ?;

-- name: CreateOrganizer :execresult
INSERT INTO organizers (id, organization_id, bank_name, bank_account, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);