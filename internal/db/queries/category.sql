-- name: AddCategory :execresult
INSERT INTO categories (id, category_name)
VALUES (?, ?);

-- name: UserCategoryExists :one
SELECT COUNT(1) from user_categories
WHERE category_id = ? AND user_id = ?;

-- name: CreateUserCategory :execresult
INSERT INTO user_categories (category_id, user_id)
VALUES (?, ?);