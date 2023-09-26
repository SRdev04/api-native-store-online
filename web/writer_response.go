package web

type WriterResponse struct {
	Id_Products int    `json:"id_products"`
	User_Id     int    `json:"user_id"`
	Name        string `json:"name"`
	Price_Carts int    `json:"price_carts"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
