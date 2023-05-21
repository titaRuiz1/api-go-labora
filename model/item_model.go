package model

type Item struct {
	ID   int    `json:"id"`
	CustomerName string `json:"customerName"`
	OrderDate string `json:"orderDate"`
	Product string `json:"product"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
}

var Items = []Item{
	{ID:1, CustomerName: "Item 1"},
	{ID:2, CustomerName: "Item 2"},
	{ID:3, CustomerName: "Item 3"},
	{ID:4, CustomerName: "Item 4"},
	{ID:5, CustomerName: "Item 5"},
	{ID:6, CustomerName: "Item 6"},
	{ID:7, CustomerName: "Item 7"},
	{ID:8, CustomerName: "Item 8"},
	{ID:9, CustomerName: "Item 9"},
	{ID:10,  CustomerName:"Item 10"},
	{ID:11,  CustomerName:"Item 11"},
	{ID:12,  CustomerName:"Item 12"},
	{ID:13,  CustomerName:"Item 13"},
	{ID:14,  CustomerName:"Item 44"},
	{ID:15,  CustomerName:"Item 15"},
	{ID:16,  CustomerName:"Item 16"},
	{ID:17,  CustomerName:"Item 17"},
	{ID:18,  CustomerName:"Item 18"},
	{ID:19,  CustomerName:"Item 19"},
	{ID:20,  CustomerName:"Item 20"},
	
}

type ItemDetails struct {
Item
Details string `json:"details"`
}