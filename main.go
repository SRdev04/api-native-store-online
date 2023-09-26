package main

import (
	"log"
	"net/http"

	"github.com/SRdev04/api-native-store-online/config"
	"github.com/SRdev04/api-native-store-online/controller"
	"github.com/SRdev04/api-native-store-online/execption"
	"github.com/SRdev04/api-native-store-online/middleware"
	"github.com/SRdev04/api-native-store-online/repository"
	"github.com/SRdev04/api-native-store-online/service"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := config.ConnectionSql()
	validate := validator.New()

	productsRepository := repository.NewProductsRepository()
	productService := service.NewProductsService(productsRepository, db, validate)
	productsController := controller.NewProductsController(productService)

	cartsRepository := repository.NewCartsRepository()
	cartsService := service.NewServiceCarts(cartsRepository, db, validate)
	cartsController := controller.NewControllerCarts(cartsService)

	shipRepository := repository.NewShippingRepository()
	shippingService := service.NewShipService(shipRepository, db, validate)
	shippingController := controller.NewControllerShipping(shippingService)

	detailRepository := repository.NewDetailRepository()
	detailService := service.NewDetailService(detailRepository, db, validate)
	detailController := controller.NewCtrlDetail(detailService)

	userRepository := repository.NewUsersRepository()
	userService := service.NewUsersService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := httprouter.New()
	router.PanicHandler = execption.ErrorHandler

	router.GET("/api/products", productsController.FindAll)
	router.GET("/api/products/:productsId", productsController.FindById)
	router.POST("/api/products", productsController.Create)
	router.PUT("/api/products/:productsId", productsController.Update)
	router.DELETE("/api/products/:productsId", productsController.Delete)

	router.GET("/api/carts", cartsController.FindAll)
	router.POST("/api/carts", cartsController.InsCarts)
	router.POST("/api/products/:productsId", cartsController.InsCarts)
	router.GET("/api/carts/:productsId", cartsController.FindIdCarts)
	router.PUT("/api/carts/:productsId", cartsController.EditQtyId)
	router.DELETE("/api/carts/:productsId", cartsController.DeleteCarts)

	router.POST("/api/shipping", shippingController.ShipInsert)
	router.GET("/api/shipping", shippingController.ShipAll)

	router.POST("/api/orders", detailController.CtrlInsDetail)
	router.GET("/api/orders", detailController.CtrlAllDetail)

	router.POST("/api/user/register", userController.Register)
	router.POST("/api/user", userController.Login)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	log.Println("Running Port:3000")

	server.ListenAndServe()

}
