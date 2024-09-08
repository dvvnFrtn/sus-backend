-- name: AddCategory :execresult
INSERT INTO categories (id, category_name, group_id)
VALUES (?, ?, ?);

-- name: AddCategoryGroup :execresult
INSERT INTO category_groups (group_name)
VALUES (?);

-- name: CategoryExists :one
SELECT COUNT(1) FROM categories WHERE category_name = ? AND group_id = ?;

-- name: CategoryGroupExists :one
SELECT COUNT(1) FROM category_groups WHERE group_name = ?;

-- name: GetCategoryGroupIDByName :one
SELECT id FROM category_groups WHERE group_name = ?;

-- name: UserCategoryExists :one
SELECT COUNT(1) FROM user_categories
WHERE category_id = ? AND user_id = ?;

-- name: CreateUserCategory :execresult
INSERT INTO user_categories (category_id, user_id)
VALUES (?, ?);

-- name: GetCategories :many
SELECT * FROM categories;

-- name: GetCategoriesForUser :many
SELECT categories.id, categories.created_at, category_name, group_id FROM categories
LEFT JOIN user_categories ON categories.id = category_id
WHERE user_id = ?;
