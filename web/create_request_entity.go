package web

type RequestCreateWeb struct {
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Price       int    `validate:"required,min=1 ,max=100" json:"price"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
