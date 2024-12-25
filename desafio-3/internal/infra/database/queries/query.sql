-- name: ListOrders :many
SELECT * FROM orders;

-- name: GetCategory :one
SELECT * FROM orders 
WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO orders (id, name, description) 
VALUES (?,?,?);

-- name: UpdateCategory :exec
UPDATE orders SET name = ?, description = ? 
WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM orders WHERE id = ?;

-- -- name: CreateCourse :exec
-- INSERT INTO courses (id, name, description, category_id, price)
-- VALUES (?,?,?,?,?);

-- -- name: ListCourses :many
-- SELECT c.*, ca.name as category_name 
-- FROM courses c JOIN orders ca ON c.category_id = ca.id;