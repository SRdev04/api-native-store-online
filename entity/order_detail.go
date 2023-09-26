package entity

type Products struct {
	Id_Products int
	Price       int
	Quantity    int
}

type Orders_Detail struct {
	Id_Orders int
	Id_User   int
	Product   []Products
}
