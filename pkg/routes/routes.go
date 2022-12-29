package routes

import (
	"github.com/gorilla/mux"
	"github.com/sanjana-jadhav-searce/shopping-catlog/pkg/controllers"
)

var RegisterShoppingCart = func(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{name}", controllers.GetProductByName).Methods("GET")
	router.HandleFunc("/updateproduct", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/deleteproduct", controllers.DeleteProduct).Methods("DELETE")

	router.HandleFunc("/createinventory", controllers.CreateInventoryProduct).Methods("POST")
	router.HandleFunc("/productinventory/{name}", controllers.GetProductByNameInInventory).Methods("GET")
	router.HandleFunc("/productsinventory", controllers.GetProductsInventory).Methods("GET")
	router.HandleFunc("/updateproductinventory", controllers.UpdateProductInventory).Methods("PUT")
	router.HandleFunc("/deleteproductinventory", controllers.DeleteProductInventory).Methods("DELETE")

	router.HandleFunc("/createcategory", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/namecategory/{name}", controllers.GetCategoryByName).Methods("GET")
	router.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/deletecategory", controllers.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/createcart", controllers.CreateCart).Methods("POST")
	router.HandleFunc("/additemstocart", controllers.AddItemsToCart).Methods("POST")
	router.HandleFunc("/additemtocart", controllers.AddItemToCart).Methods("POST")
	router.HandleFunc("/getcart", controllers.GetCart).Methods("GET")
	router.HandleFunc("/deletecart", controllers.DeleteCart).Methods("DELETE")

}
