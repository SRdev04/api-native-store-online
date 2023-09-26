package web

type Products struct {
	Id_Products int `json:"id_products"`
	Price       int `json:"price"`
	Quantity    int `json:"quantity"`
}

type DetailRequest struct {
	Id_Orders int `json:"id_orders"`
	Id_User   int `json:"id_user"`
	Product   []Products
}
