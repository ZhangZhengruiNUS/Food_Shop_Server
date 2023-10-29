// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"
)

type Querier interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	DeleteProduct(ctx context.Context, productID int64) error
	GetProductCount(ctx context.Context) (int64, error)
	GetProductCountByOwner(ctx context.Context, shopOwnerID int64) (int64, error)
	GetProductList(ctx context.Context, arg GetProductListParams) ([]GetProductListRow, error)
	GetProductListByOwner(ctx context.Context, arg GetProductListByOwnerParams) ([]GetProductListByOwnerRow, error)
}

var _ Querier = (*Queries)(nil)
