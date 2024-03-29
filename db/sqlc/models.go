// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"time"
)

type Product struct {
	ProductID     int64     `json:"productId"`
	ShopOwnerName string    `json:"shopOwnerName"`
	PicPath       string    `json:"picPath"`
	Describe      string    `json:"describe"`
	Price         float64   `json:"price"`
	Quantity      int32     `json:"quantity"`
	ExpireTime    time.Time `json:"expireTime"`
	CreateTime    time.Time `json:"createTime"`
}
