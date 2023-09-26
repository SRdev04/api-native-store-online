package helper

import (
	"github.com/SRdev04/api-native-store-online/entity"

	"github.com/SRdev04/api-native-store-online/web"
)

func ConvertTo(carts entity.ResultCarts) web.WriterResponse {
	voka := web.WriterResponse{
		Id_Products: carts.Id_Products,
		User_Id:     carts.User_Id,
		Name:        carts.Name,
		Price_Carts: carts.Price_Carts,
		Description: carts.Description,
		Quantity:    carts.Quantity,
	}

	return voka

}

func ForGetAll(carts []entity.ResultCarts) []web.WriterResponse {
	var cartsRespon []web.WriterResponse

	for _, result := range carts {

		cartsRespon = append(cartsRespon, ConvertTo(result))

	}

	return cartsRespon

}

func ConvertInsertCart(carts entity.ResultCarts) web.WebCarts {
	tampung := web.WebCarts{
		Id:          carts.Id,
		User_Id:     carts.User_Id,
		Id_Products: carts.Id_Products,
		Price_Carts: carts.Price_Carts,
		Quantity:    carts.Quantity,
	}
	return tampung
}

func ConvertShipping(ship entity.Shipping) web.ShipingRespon {
	convertShip := web.ShipingRespon{
		User_Id:   ship.User_Id,
		Kota:      ship.Kota,
		Kecamatan: ship.Kecamatan,
		Kelurahan: ship.Kelurahan,
		Jalan:     ship.Jalan,
		Wa:        ship.Wa,
	}
	return convertShip

}
func ConvertShipSlice(ship []entity.Shipping) []web.ShipingRespon {
	var web []web.ShipingRespon

	for _, result := range ship {
		web = append(web, ConvertShipping(result))
	}
	return web
}

func ConvToSlice(req web.DetailRequest) entity.Orders_Detail {
	var order entity.Orders_Detail

	try := order.Product

	for _, result := range req.Product {
		try = append(try, entity.Products(result))
	}

	conv := entity.Orders_Detail{
		Id_Orders: req.Id_Orders,
		Id_User:   req.Id_User,
		Product:   try,
	}
	return conv
}

func ConvTodetail(req web.DetailRequest) entity.Orders_Detail {

	var tampung entity.Orders_Detail
	try := tampung.Product

	for _, result := range req.Product {
		try = append(try, entity.Products(result))
	}
	return tampung
}

func ConvToRes(order entity.Orders_Detail) web.DetailRespon {

	result := web.DetailRespon{
		Id_Orders: order.Id_Orders,
		Id_User:   order.Id_User,
	}
	resultProduct := result.Product

	for _, rows := range order.Product {
		resultProduct = append(resultProduct, web.Respon(rows))
	}
	result2 := web.DetailRespon{
		Id_Orders: order.Id_Orders,
		Id_User:   order.Id_User,
		Product:   resultProduct,
	}

	return result2
}

func ConvertProduct(product entity.ProductsTabel) web.WritertWeb {
	return web.WritertWeb{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
	}

}

func ForFindAll(product []entity.ProductsTabel) []web.WritertWeb {
	var productsResponse []web.WritertWeb

	for _, result := range product {
		productsResponse = append(productsResponse, ConvertProduct(result))

	}
	return productsResponse

}

func UserConv(user entity.Users) web.UsersRespon {
	return web.UsersRespon{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

}

//func ConvToResponDetail(order entity.AllOrders) web.ResponDetail {
//	result := web.ResponDetail{
//		Id_Orders: order.Id_Orders,
//		Id_User:   order.Id_User,
//	}
//
//	conv := result.Products
//
//	for _, rows := range order.Products {
//		conv = append(conv, web.DetailProduct(rows))
//	}
//
//	allResult := web.ResponDetail{
//		Id_Orders: order.Id_Orders,
//		Id_User:   order.Id_User,
//		Products:  conv,
//	}
//	return allResult
//
//}
//
//func ConvAllToDetail(orders []entity.AllOrders) []web.ResponDetail {
//	var respons []web.ResponDetail
//
//	for _, rows := range orders {
//
//		respons = append(respons, ConvToResponDetail(rows))
//	}
//
//	return respons
//}
//

func ConvertAll(order entity.AllOrders) web.ResponDetail {
	conv := web.ResponDetail{
		Id_Orders:   order.Id_Orders,
		Id_User:     order.Id_User,
		Id_Products: order.Id_Products,
		Name:        order.Name,
		Description: order.Description,
		Quantity:    order.Quantity,
		Price:       order.Price,
	}
	return conv

}

func ConvertAllSlice(order []entity.AllOrders) []web.ResponDetail {
	var conv []web.ResponDetail

	for _, rows := range order {

		conv = append(conv, web.ResponDetail(rows))

	}
	return conv
}
