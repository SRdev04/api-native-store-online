package web

type WritertWeb struct {
	Id          int    `json:"id"`
	Name        string `validate:"required,min1"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
