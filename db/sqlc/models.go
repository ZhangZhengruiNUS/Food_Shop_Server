// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"time"
)

type Product struct {
	ProductID   int64     `json:"productId"`
	ShopOwnerID int64     `json:"shopOwnerId"`
	PicPath     string    `json:"picPath"`
	Describe    string    `json:"describe"`
	Price       int32     `json:"price"`
	Quantity    int32     `json:"quantity"`
	CreateTime  time.Time `json:"createTime"`
}
