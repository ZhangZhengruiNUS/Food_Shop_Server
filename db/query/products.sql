-- name: GetProductCount :one
SELECT COUNT(*)
FROM products LIMIT 1;

-- name: GetProductCountByOwner :one
SELECT COUNT(*)
FROM products
WHERE shop_owner_id = $1 LIMIT 1;

-- name: GetProductList :many
SELECT product_id, product_name, pic_path
FROM products
LIMIT $2
OFFSET (($1 - 1) * $2);

-- name: GetProductListByOwner :many
SELECT product_id, product_name, pic_path
FROM products
WHERE shop_owner_id = $1
LIMIT $3
OFFSET (($2 - 1) * $3);
