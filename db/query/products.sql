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
WHERE product_id NOT IN (SELECT *
                         FROM products LIMIT (($1 - 1) * $2)
    ) LIMIT $2;

-- name: GetProductListByOwner :many
SELECT product_id, product_name, pic_path
FROM products
WHERE shop_owner_id = $1
  AND product_id NOT IN (SELECT *
                         FROM products
                         WHERE shop_owner_id = $1
                         LIMIT (($2 - 1) * $3)
    ) LIMIT $3;
