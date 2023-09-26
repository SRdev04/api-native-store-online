package web

type DetailProduct struct {
	Id_Products int
	Name        string
	Description string
	Quantity    int
	Price       int
}

type ResponDetail struct {
	Id_Orders   int
	Id_User     int
	Id_Products int
	Name        string
	Description string
	Quantity    int
	Price       int
}
