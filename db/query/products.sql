-- name: GetProductCount :one
SELECT COUNT(*)
FROM products LIMIT 1;

-- name: GetProductCountByOwner :one
SELECT COUNT(*)
FROM products
WHERE shop_owner_id = $1 LIMIT 1;

-- name: GetProductList :many
SELECT product_id, describe, pic_path
FROM products
LIMIT $2
OFFSET (($1 - 1) * $2);

-- name: GetProductListByOwner :many
SELECT product_id, describe, pic_path
FROM products
WHERE shop_owner_id = $1
LIMIT $3
OFFSET (($2 - 1) * $3);

-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id =$1;

-- name: CreateProduct :one
INSERT INTO products(
  shop_owner_id,
  pic_path,
  describe,
  price,
  quantity
  ) VALUES (
  $1, $2, $3, $4, $5
  )
  RETURNING *;
)
