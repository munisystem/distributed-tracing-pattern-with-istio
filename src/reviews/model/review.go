package model

type Review struct {
	Id          int64  `json:"id"`
	CustomerId  int64  `json:"customer_id"`
	ItemId      int64  `json:"item_id"`
	Rating      int64  `json:"rating"`
	Description string `json:"description"`
}
