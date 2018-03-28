package model

type Review struct {
	ItemName    string `json:"item_name,omitempty"`
	Rating      int64  `json:"rating,omitempty"`
	Description string `json:"description,omitempty"`
	Message     string `json:"message,omitempty"`
}
