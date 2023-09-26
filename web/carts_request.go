package web

type RequestCarts struct {
	User_Id     int `json:"user_id"`
	Id_Products int `json:"id_products"`
	Price_Carts int `json:"price_carts"`
	Quantity    int `json:"quantity"`
}
