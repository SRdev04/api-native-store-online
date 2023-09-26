package web

type CartRequest struct {
	Cartid    string
	Productid string
	Info      map[string]interface{}
}

type CartRespon struct {
	Cartid    string
	Productid string
	Info      map[string]interface{}
}
