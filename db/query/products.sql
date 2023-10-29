-- name: GetProductCount :one
SELECT COUNT(*) FROM products
LIMIT 1;

-- name: GetProductCountByOwner :one
SELECT COUNT(*) FROM products
WHERE shop_owner_name = $1 LIMIT 1;

-- name: GetProductList :many
SELECT product_id, describe, pic_path
FROM products
LIMIT sqlc.arg(pageSize)::int
OFFSET ((sqlc.arg(page)::int - 1) * sqlc.arg(pageSize)::int);

-- name: GetProductListByOwner :many
SELECT product_id, describe, pic_path
FROM products
WHERE shop_owner_name = $1
LIMIT sqlc.arg(pageSize)::int
OFFSET ((sqlc.arg(page)::int - 1) * sqlc.arg(pageSize)::int);

-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1;

-- name: CreateProduct :one
INSERT INTO products (
  shop_owner_name,
  pic_path,
  describe,
  price,
  quantity
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;
