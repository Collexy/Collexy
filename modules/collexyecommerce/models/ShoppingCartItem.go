package models

type ShoppingCartItem struct {
	Id int `json:"id,omitempty"`
	Quantity int `json:"quantity,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}