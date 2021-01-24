package _type

import (
//"go.mongodb.org/mongo-driver/bson/primitive"
//"time"
)

//confused
type GetListInput struct {
	AccountId string `json:"accountId"`
	Search    string `json:"search"`
	Start     int    `json:"start"`
	Limit     int    `json:"limit"`
}
type SortParam struct {
	SortBy string // bson field name
	Type   int    // -1 or 1
}

type GetListFilterInput struct {
	FetchMode       string            `json:"fetchMode"`
	FilterParams    string            `json:"filterParams"`
	FilterParamsMap map[string]string `json:"filterParamsMap"`
}

type SearchFilterAndPaginationInput struct {
	Search       string            `json:"search"`
	FilterParams map[string]string `json:"filterParams"`
	Page         int               `json:"page"`
	Pagelen      int               `default:"20" json:"pagelen"`
	Sort         string            `json:"sort"`
	Desc         bool              `json:"desc"`
}
