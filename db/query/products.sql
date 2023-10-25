-- name: GetProductCount :one
SELECT COUNT(*) FROM products
LIMIT 1;

-- name: GetProductCountByOwner :one
SELECT COUNT(*) FROM products
WHERE shop_owner_id = $1 LIMIT 1;