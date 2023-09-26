package web

type RequestUpdateWeb struct {
	Id          int    `json:"id"`
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Price       int    `validate:"required,min=1,max=100" json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
